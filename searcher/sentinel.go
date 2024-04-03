package searcher

import (
	"context"
	"fmt"
	"gar/ent"
	"gar/registry"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"sync/atomic"
	"time"
)

var (
	ErrNoAliveWorker = errors.New("no alive worker")
)

type Sentinel struct {
	hub   *registry.HubProxy
	pools sync.Map
}

func NewSentinel(endpoints registry.EtcdEndpoints) *Sentinel {
	return &Sentinel{
		hub:   registry.NewHubProxy(endpoints, 10, 100),
		pools: sync.Map{},
	}
}

func (s *Sentinel) Conn(endpoint string) *grpc.ClientConn {
	if value, has := s.pools.Load(endpoint); has {
		conn := value.(*grpc.ClientConn)
		stat := conn.GetState()

		if stat == connectivity.TransientFailure || stat == connectivity.Shutdown {
			slog.Warn("connection status to endpoint {} is {}", endpoint, stat)
			_ = conn.Close()
		} else {
			return conn
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		slog.Error("dial {} failed: {}", endpoint, err)
		return nil
	}

	slog.Info("connected to grpc server {}", endpoint)
	s.pools.Store(endpoint, conn)

	return conn
}

func (s *Sentinel) Add(doc *ent.Document) (int, error) {
	endpoint := s.hub.Endpoint(SearchService)

	if len(endpoint) == 0 {
		return 0, ErrNoAliveWorker
	}

	conn := s.Conn(endpoint)

	if conn == nil {
		return 0, fmt.Errorf("connect worker %s failed", endpoint)
	}

	client := NewSearchServiceClient(conn)
	affected, err := client.Add(context.Background(), doc)

	if err != nil {
		return 0, err
	}

	slog.Info("add {} docs to {}", affected.Count, endpoint)

	return int(affected.Count), nil
}

func (s *Sentinel) Delete(docId string) int {
	endpoints := s.hub.Endpoints(SearchService)

	if len(endpoints) == 0 {
		return 0
	}

	var n int32
	wg := sync.WaitGroup{}
	wg.Add(len(endpoints))

	for _, endpoint := range endpoints {
		go func(endpoint string) {
			defer wg.Done()

			conn := s.Conn(endpoint)

			if conn != nil {
				client := NewSearchServiceClient(conn)
				affected, err := client.Delete(context.Background(), &DocId{Id: docId})

				if err == nil {
					atomic.AddInt32(&n, affected.Count)
					slog.Info("delete {} docs from worker {}", affected.Count, endpoint)
				} else {
					slog.Error("delete doc {} from worker {} failed: {}", docId, endpoint, err)
				}
			}
		}(endpoint)
	}

	wg.Wait()

	return int(atomic.LoadInt32(&n))
}

func (s *Sentinel) Search(q *ent.TermQuery, onFlag uint64, offFlag uint64, orFlags []uint64) []*ent.Document {
	endpoints := s.hub.Endpoints(SearchService)

	if len(endpoints) == 0 {
		return nil
	}

	docs := make([]*ent.Document, 0, 1000)
	docChan := make(chan *ent.Document, 1000)

	wg := sync.WaitGroup{}
	wg.Add(len(endpoints))

	for _, endpoint := range endpoints {
		go func(endpoint string) {
			defer wg.Done()
			conn := s.Conn(endpoint)

			if conn != nil {
				client := NewSearchServiceClient(conn)
				r, err := client.Search(context.Background(), &SearchRequest{
					Query:   q,
					OnFlag:  onFlag,
					OffFlag: offFlag,
					OrFlags: orFlags,
				})

				if err == nil {
					if len(r.Documents) > 0 {
						slog.Info("matched {} docs from worker {}", len(r.Documents), endpoint)
						for _, doc := range r.Documents {
							docChan <- doc
						}
					}
				} else {
					slog.Warn("search from cluster failed {}", err)
				}
			}
		}(endpoint)
	}

	finish := make(chan struct{})

	go func() {
		for {
			doc, ok := <-docChan
			if !ok {
				break
			}

			docs = append(docs, doc)
		}

		finish <- struct{}{}
	}()

	wg.Wait()
	close(docChan)
	<-finish

	return docs
}

func (s *Sentinel) Count() int {
	endpoints := s.hub.Endpoints(SearchService)

	if len(endpoints) == 0 {
		return 0
	}

	var n int32
	wg := sync.WaitGroup{}
	wg.Add(len(endpoints))

	for _, endpoint := range endpoints {
		go func(endpoint string) {
			defer wg.Done()
			conn := s.Conn(endpoint)

			if conn != nil {
				client := NewSearchServiceClient(conn)
				affected, err := client.Count(context.Background(), &CountRequest{})

				if err == nil {
					if affected.Count > 0 {
						atomic.AddInt32(&n, affected.Count)
						slog.Info("worker {} have {} documents", endpoint, affected.Count)
					}
				} else {
					slog.Warn("get doc count from worker {} failed: {}", endpoint, err)
				}
			}
		}(endpoint)
	}

	wg.Wait()

	return int(atomic.LoadInt32(&n))
}

func (s *Sentinel) Close() error {
	var err error

	s.pools.Range(func(_, value any) bool {
		conn := value.(*grpc.ClientConn)
		err = conn.Close()
		return true
	})

	if err != nil {
		return err
	}

	return s.hub.Close()
}
