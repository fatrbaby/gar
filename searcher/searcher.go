package searcher

import (
	"context"
	"fmt"
	"gar/ent"
	"gar/registry"
	"gar/shortcut"
	"github.com/pkg/errors"
	"time"
)

const SearchService = "search-service"

type Searcher struct {
	Indexer *Indexer
	hub     *registry.HubProxy
	addr    string
}

func NewSearcher() *Searcher {
	return &Searcher{}
}

func (s *Searcher) Setup(docEstimate int, storagePath string) error {
	s.Indexer = NewIndexer()
	return s.Indexer.Setup(docEstimate, storagePath)
}

func (s *Searcher) Online(endpoints registry.EtcdEndpoints, port int) error {
	if len(endpoints) == 0 {
		return nil
	}

	if port < 1024 {
		return errors.New("invalid listen port, must more than 1024")
	}

	localIp, err := shortcut.LocalIp()

	if err != nil {
		panic(err)
	}

	s.addr = fmt.Sprintf("%s:%d", localIp, port)
	var heartbeat int64 = 3
	hub := registry.NewHubProxy(endpoints, heartbeat, 1000)
	s.hub = hub
	leaseID, err := hub.Register(SearchService, s.addr, 0)

	if err != nil {
		return err
	}

	go func() {
		for {
			_, _ = hub.Register(SearchService, s.addr, leaseID)
			time.Sleep(time.Duration(heartbeat)*time.Second - 100*time.Millisecond)
		}
	}()

	return nil
}

func (s *Searcher) Search(_ context.Context, r *SearchRequest) (*SearchResult, error) {
	docs := s.Indexer.Search(r.Query, r.OnFlag, r.OffFlag, r.OrFlags)

	return &SearchResult{Documents: docs}, nil
}

func (s *Searcher) Add(_ context.Context, doc *ent.Document) (*AffectedCount, error) {
	n, err := s.Indexer.Add(doc)

	return &AffectedCount{Count: int32(n)}, err
}

func (s *Searcher) Delete(_ context.Context, docId *DocId) (*AffectedCount, error) {
	return &AffectedCount{
		Count: int32(s.Indexer.Delete(docId.Id)),
	}, nil
}

func (s *Searcher) Count(ctx context.Context, _ *CountRequest) (*AffectedCount, error) {
	return &AffectedCount{Count: int32(s.Indexer.Count())}, nil
}

func (s *Searcher) LoadFromStorage() int {
	return s.Indexer.LoadFromFile()
}

func (s *Searcher) Close() error {
	if s.hub != nil {
		return s.hub.Offline(SearchService, s.addr)
	}

	return s.Indexer.Close()
}

func (s *Searcher) mustEmbedUnimplementedSearchServiceServer() {
}
