package main

import (
	"gar/app/search"
)

func main() {
	options := &search.WorkerOptions{
		Host:          "127.0.0.1",
		Port:          6123,
		Workdir:       "/Users/fatrbaby/workspaces/golang/gar/var/",
		Number:        1,
		RebuildIndex:  false,
		EtcdEndpoints: []string{"127.0.0.1:2379"},
	}

	worker := search.NewSearcherWorker(options)
	worker.Run()
}
