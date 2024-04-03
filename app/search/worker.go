package search

import (
	"gar/app/data"
	"gar/searcher"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Worker struct {
	ID         int
	options    *WorkerOptions
	service    *searcher.Searcher
	datasource data.Source
}

func NewSearcherWorker(options *WorkerOptions) *Worker {
	return &Worker{
		options: options,
		service: searcher.NewSearcher(),
	}
}

func (w *Worker) Run() {
	go w.teardown()
	w.startup()
}

func (w *Worker) WithDatasource(ds data.Source) *Worker {
	w.datasource = ds

	return w
}

func (w *Worker) startup() {
	err := w.service.Setup(100000, w.options.StoragePath())

	if err != nil {
		panic(err)
	}

	if w.options.RebuildIndex {
		if w.datasource != nil {
			w.datasource.BuildIndexes(w.service.Indexer, w.options.Total, w.options.ID)
		} else {
			slog.Warn("No datasource set when rebuild indexes")
		}
	} else {
		w.service.LoadFromStorage()
	}

	slog.Info("register searcher worker")
	err = w.service.Online(w.options.EtcdEndpoints, w.options.Port)

	if err != nil {
		slog.Error("register searcher worker failed: {}", err)
	} else {
		slog.Info("register searcher worker success")
	}

	listener, err := net.Listen("tcp", w.options.Addr())

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	searcher.RegisterSearchServiceServer(server, w.service)

	if err = server.Serve(listener); err != nil {
		_ = w.service.Close()
		slog.Error("start grpc server on port :{} failed: {}", w.options.Port, err)
	}
}

func (w *Worker) teardown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := w.service.Close(); err != nil {
		slog.Warn("quit search worker failed: {}", err)
	} else {
		slog.Info("quit search worker success")
	}

	os.Exit(0)
}
