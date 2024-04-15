package main

import (
	"gar/app/data"
	"gar/app/handler"
	"gar/app/storage"
	"gar/shortcut"
	"github.com/gofiber/fiber/v2"
	"log"
	"path"
)

func main() {
	w := createStandaloneWorker()

	go func() {
		w.Run()
	}()

	app := fiber.New(fiber.Config{
		AppName: "Gar",
	})

	h := handler.New(w.Indexer())

	app.Static("/", "./public")

	api := app.Group("/api")
	api.Get("/categories", h.Categories)
	api.Post("/search", h.Search)

	log.Fatalln(app.Listen(":5321"))
}

func createStandaloneWorker() storage.Worker {
	cwd := shortcut.CurrentPath()
	workdir := path.Join(cwd, "var")

	worker := storage.NewStandaloneWorker(path.Join(workdir, "standalone"), false)
	worker.WithDatasource(data.NewCsvSource(workdir + "/bili_video.csv"))

	return worker
}

func createDistributedWorker() storage.Worker {
	cwd := shortcut.CurrentPath()
	workdir := path.Join(cwd, "var")

	options := &storage.WorkerOptions{
		ID:            0,
		Host:          "127.0.0.1",
		Port:          6123,
		Workdir:       path.Join(workdir, "db"),
		BuildIndexes:  false,
		EtcdEndpoints: []string{"127.0.0.1:2379"},
		NumWorkers:    1,
	}

	worker := storage.NewDistributedWorker(options)
	worker.WithDatasource(data.NewCsvSource(workdir + "/bili_video.csv"))

	return worker
}
