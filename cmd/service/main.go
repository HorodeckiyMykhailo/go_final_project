package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/HorodeckiyMykhailo/go_final_project/internal/db"
	"github.com/HorodeckiyMykhailo/go_final_project/internal/handler"
	"github.com/HorodeckiyMykhailo/go_final_project/internal/repository"
	"github.com/go-chi/chi"
)

const defaultPort = ":7540"
const webDir = "./web"

func migration(rep *repository.Repository){
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath),"scheduler.db")
	_, err = os.Stat(dbFile)

	var install bool 
	if err != nil {
		install = true
	}
	
	if install {
		if err := rep.CreateScheduler();err != nil{
			log.Fatal(err)
		}
	}
}




func main() {
	fmt.Println("Запуск сервера")

	db := db.New()
	repo := repository.New(db)
	migration(repo)
	handler := handler.New(repo)

	r := chi.NewRouter()
	r.Handle("/*",http.FileServer(http.Dir(webDir)))

	r.Post("/api/task/done",handler.Done)
	r.Delete("/api/task",handler.Delete)
	r.Put("/api/task",handler.UpdateTask)
	r.Get("/api/task", handler.GetTaskById)
	r.Get("/api/tasks",handler.GetTasks)
	r.Post("/api/task",handler.AddTask)
	r.Get("/api/nextdate",handler.NextDateHandler)

	if err := http.ListenAndServe(defaultPort,r); err != nil {
		log.Fatal(err)
	}
}


