package main

import (
	"gar/app/data"
	"gar/app/search"
	"gar/shortcut"
	"path"
)

func main() {
	cwd := shortcut.CurrentPath()
	workdir := path.Join(cwd, "var")

	options := &search.WorkerOptions{
		ID:            0,
		Host:          "127.0.0.1",
		Port:          6123,
		Workdir:       workdir,
		RebuildIndex:  false,
		EtcdEndpoints: []string{"127.0.0.1:2379"},
		NumWorkers:    1,
	}

	worker := search.NewSearcherWorker(options)
	worker.WithDatasource(data.NewCsvSource(workdir + "/bili_video.csv"))
	worker.Run()
}
