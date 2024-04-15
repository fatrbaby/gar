package storage

import (
	"gar/app/data"
	"gar/searcher"
	"os"
	"os/signal"
	"syscall"
)

type StandaloneWorker struct {
	workdir        string
	rebuildIndexes bool
	datasource     data.Source
	indexer        *searcher.Indexer
}

func NewStandaloneWorker(workdir string, rebuild bool) Worker {
	return &StandaloneWorker{
		workdir:        workdir,
		rebuildIndexes: rebuild,
		indexer:        searcher.NewIndexer(),
	}
}

func (s *StandaloneWorker) WithDatasource(source data.Source) Worker {
	s.datasource = source

	return s
}

func (s *StandaloneWorker) Run() {
	go s.teardown()
	s.startup()
}

func (s *StandaloneWorker) Indexer() *searcher.Indexer {
	return s.indexer
}

func (s *StandaloneWorker) startup() {
	err := s.indexer.Setup(100000, s.workdir)

	if err != nil {
		panic(err)
	}

	if s.rebuildIndexes {
		s.datasource.BuildIndexes(s.indexer, 0, 0)
	} else {
		s.indexer.LoadFromFile()
	}
}

func (s *StandaloneWorker) teardown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.indexer.Close() //接收到kill信号时关闭索引
	os.Exit(0)
}
