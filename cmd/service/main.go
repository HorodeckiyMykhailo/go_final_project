package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/HorodeckiyMykhailo/go_final_project/internal/db"
	"github.com/HorodeckiyMykhailo/go_final_project/internal/handler"
	"github.com/HorodeckiyMykhailo/go_final_project/internal/repository"
	"github.com/go-chi/chi"
)

const defaultPort = ":7540"
const webDir = "./web"



func main() {
	fmt.Println("Запуск сервера")
	log.Printf("Сервер запущен на порту: %s\n", defaultPort)

	data := db.New()
	repo := repository.New(data)
	db.Migration(repo)
	handlers := handler.New(repo)

	r := chi.NewRouter()
	r.Handle("/*",http.FileServer(http.Dir(webDir)))

	r.Post("/api/task/done",handlers.Done)
	r.Delete("/api/task",handlers.Delete)
	r.Put("/api/task",handlers.UpdateTask)
	r.Get("/api/task", handlers.GetTaskById)
	r.Get("/api/tasks",handlers.GetTasks)
	r.Post("/api/task",handlers.AddTask)
	r.Get("/api/nextdate",handlers.NextDateHandler)

	if err := http.ListenAndServe(defaultPort,r); err != nil {
		log.Fatal(err)
	}
}


