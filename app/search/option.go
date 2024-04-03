package search

import (
	"fmt"
	"gar/registry"
	"path"
)

type WorkerOptions struct {
	ID            int
	Host          string
	Port          int
	Workdir       string
	RebuildIndex  bool
	EtcdEndpoints registry.EtcdEndpoints
	Total         int
}

func (o *WorkerOptions) Addr() string {
	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}

func (o *WorkerOptions) StoragePath() string {
	return path.Join(o.Workdir, fmt.Sprintf("%s%d", "_part", o.ID))
}
