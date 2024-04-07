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
		ID:            1,
		Host:          "127.0.0.1",
		Port:          6123,
		Workdir:       workdir,
		RebuildIndex:  true,
		EtcdEndpoints: []string{"127.0.0.1:2379"},
		Total:         1,
	}

	worker := search.NewSearcherWorker(options)
	worker.WithDatasource(data.NewCsvSource(workdir + "/bili_video.csv"))
	worker.Run()
}
