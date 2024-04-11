package search

import (
	"gar/app/ent"
	"gar/app/search/bus"
	"gar/app/search/filterer"
	"gar/app/search/recaller"
	"golang.org/x/exp/maps"
	"log/slog"
	"reflect"
	"sync"
	"time"
)

type Recaller interface {
	Recall(*bus.Context) []*ent.BiliBiliVideo
}

type Filterer interface {
	Apply(*bus.Context)
}

type VideoSearcher struct {
	Recallers []Recaller
	Filterers []Filterer
}

func (s *VideoSearcher) WithRecaller(recallers ...Recaller) *VideoSearcher {
	s.Recallers = append(s.Recallers, recallers...)
	return s
}

func (s *VideoSearcher) WithFilterer(filterers ...Filterer) *VideoSearcher {
	s.Filterers = append(s.Filterers, filterers...)
	return s
}

func (s *VideoSearcher) Recall(ctx *bus.Context) {
	if len(s.Recallers) == 0 {
		return
	}

	collection := make(chan *ent.BiliBiliVideo, 1000)
	wg := sync.WaitGroup{}
	wg.Add(len(s.Recallers))

	for _, recaller := range s.Recallers {
		go func(recaller Recaller) {
			defer wg.Done()
			rule := reflect.TypeOf(recaller).Name()
			result := recaller.Recall(ctx)
			slog.Info("recalled results", "count", len(result), "recaller", rule)

			for _, video := range result {
				collection <- video
			}

		}(recaller)
	}

	results := make(map[string]*ent.BiliBiliVideo, 1000)
	finished := make(chan struct{})

	go func() {
		for {
			video, ok := <-collection
			if !ok {
				break
			}
			results[video.Id] = video
		}
		finished <- struct{}{}
	}()
	wg.Wait()
	close(collection)
	<-finished

	ctx.Results = maps.Values(results)
}

func (s *VideoSearcher) Filter(c *bus.Context) {
	for _, f := range s.Filterers {
		f.Apply(c)
	}
}

func (s *VideoSearcher) Search(c *bus.Context) []*ent.BiliBiliVideo {
	t1 := time.Now()
	s.Recall(c)
	t2 := time.Now()
	slog.Info("recall done", "count", len(c.Results), "duration", t2.Sub(t1).Milliseconds())

	s.Filter(c)
	t3 := time.Now()
	slog.Info("recall done", "remain", len(c.Results), "duration", t3.Sub(t2).Milliseconds())

	return c.Results
}

func NewVideoSearcher() *VideoSearcher {
	s := &VideoSearcher{}
	s.WithRecaller(&recaller.KeywordRecaller{})
	s.WithFilterer(&filterer.ViewRangeFilterer{})

	return s
}

type AuthorVideoSearcher struct {
	VideoSearcher
}

func NewAuthOrVideoSearcher() *AuthorVideoSearcher {
	s := &AuthorVideoSearcher{}
	s.WithRecaller(&recaller.KeywordAndAuthorRecaller{})
	s.WithFilterer(&filterer.ViewRangeFilterer{})

	return s
}
