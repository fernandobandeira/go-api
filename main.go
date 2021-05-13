package main

import (
	"api/api/handlers"
	"api/entities"
	"api/repositories"
	"api/services"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// DB connection
	const dsn = "postgresql://postgres:postgres@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		os.Exit(1)
	}

	db.AutoMigrate(entities.Todo{})

	// Repositories
	todoRepository := repositories.NewTodoRepository(db)

	// Services
	todoService := services.NewTodoService(todoRepository)

	r := chi.NewRouter()

	r.Route("/v1", func(v1 chi.Router) {
		v1.Mount("/todos", handlers.TodoRouter(todoService))
	})

	srv := http.Server{
		Addr:    ":3000",
		Handler: r,

		// These values are here to make sure that the server doesn't hang
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		// This value is extremely important, it prevents us from suffering a Slowloris attack
		IdleTimeout: time.Second * 60,
	}
	srv.ListenAndServe()
}
