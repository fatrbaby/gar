package main

import (
	"gar/app/data"
	"gar/app/handler"
	"gar/app/storage"
	"gar/shortcut"
	"github.com/gofiber/fiber/v3"
	"log"
	"path"
)

func main() {
	go worker()
	
	app := fiber.New(fiber.Config{
		AppName: "Gar",
	})

	app.Static("/", "./public")

	api := app.Group("/api")
	api.Get("/categories", handler.Categories)
	api.Post("/search", handler.Search)

	log.Fatalln(app.Listen(":5321"))
}

func worker() {
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

	worker := storage.NewWorker(options)
	worker.WithDatasource(data.NewCsvSource(workdir + "/bili_video.csv"))
	worker.Run()
}
