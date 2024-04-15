package storage

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

type DistributedWorker struct {
	ID         int
	options    *WorkerOptions
	sentinel   *searcher.Sentinel
	datasource data.Source
}

func NewDistributedWorker(options *WorkerOptions) *DistributedWorker {
	return &DistributedWorker{
		options:  options,
		sentinel: searcher.NewSentinel(options.EtcdEndpoints),
	}
}

func (w *DistributedWorker) Run() {
	go w.teardown()
	w.startup()
}

func (w *DistributedWorker) WithDatasource(ds data.Source) Worker {
	w.datasource = ds

	return w
}

func (w *DistributedWorker) Indexer() *searcher.Indexer {
	return nil
}

func (w *DistributedWorker) startup() {
	var err error
	//err := w.sentinel.Setup(100000, w.options.StoragePath())

	//if err != nil {
	//panic(err)
	//}

	if w.options.BuildIndexes {
		if w.datasource != nil {
			slog.Info("start rebuild indexes")
			// w.datasource.BuildIndexes(w.sentinel.Indexer, w.options.NumWorkers, w.options.ID)
			slog.Info("rebuild indexes done")
		} else {
			slog.Warn("No datasource set when rebuild indexes")
		}
	} else {
		// w.sentinel.LoadFromStorage()
	}

	slog.Info("start register searcher worker")
	// err = w.sentinel.Online(w.options.EtcdEndpoints, w.options.Port)

	if err != nil {
		slog.Error("register searcher worker failed", "error", err)
	} else {
		slog.Info("register searcher worker done")
	}

	listener, err := net.Listen("tcp", w.options.Addr())

	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	// searcher.RegisterSearchServiceServer(server, w.sentinel)

	if err = server.Serve(listener); err != nil {
		_ = w.sentinel.Close()
		slog.Error("start grpc server failed", "port", w.options.Port, "error", err)
	}
}

func (w *DistributedWorker) teardown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("close searcher worker server...")

	if err := w.sentinel.Close(); err != nil {
		slog.Warn("quit search worker failed", "error", err)
	} else {
		slog.Info("quit search worker success")
	}

	os.Exit(0)
}
