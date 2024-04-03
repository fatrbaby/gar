package search

import (
	"gar/app/source"
	"gar/searcher"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Worker struct {
	options *WorkerOptions
	service *searcher.Searcher
}

func NewSearcherWorker(options *WorkerOptions) *Worker {
	return &Worker{
		options: options,
		service: searcher.NewSearcher(),
	}
}

func (s *Worker) Run() {
	go s.teardown()
	s.startup()
}

func (s *Worker) startup() {
	err := s.service.Setup(100000, s.options.StoragePath())

	if err != nil {
		panic(err)
	}

	if s.options.RebuildIndex {
		source.BuildIndexes(source.CsvDataSource, s.service.Indexer, 1, s.options.Number)
	} else {
		s.service.LoadFromStorage()
	}

	slog.Info("register searcher worker")
	err = s.service.Online(s.options.EtcdEndpoints, s.options.Port)

	if err != nil {
		slog.Error("register searcher worker failed: {}", err)
	} else {
		slog.Info("register searcher worker success")
	}

	listener, err := net.Listen("tcp", s.options.Addr())

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	searcher.RegisterSearchServiceServer(server, s.service)

	if err = server.Serve(listener); err != nil {
		_ = s.service.Close()
		slog.Error("start grpc server on port :{} failed: {}", s.options.Port, err)
	}
}

func (s *Worker) teardown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := s.service.Close(); err != nil {
		slog.Warn("quit search worker failed: {}", err)
	} else {
		slog.Info("quit search worker success")
	}

	os.Exit(0)
}
