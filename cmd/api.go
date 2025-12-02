package main

import (
	"myproject/internal/json"
	"myproject/internal/products"
	repo "myproject/internal/sqlc/out"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

func (app *application) mount() http.Handler{
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		json.Write(w, 200, "Hello World")
	})

	r.Route("/todos", func(r chi.Router){
		todoHandler := products.NewHandler(repo.New(app.db))

		r.Get("/", todoHandler.ListTodos)
		r.Get("/{id}", todoHandler.GetById)
	})

	return r
}

func run(application *application, h http.Handler) error{
	srv := &http.Server{
		Addr: application.config.addr,
		Handler: h,
	}
	return srv.ListenAndServe()
}
type application struct {
	config config
	db *pgx.Conn
}

type config struct {
	addr string
	dbConfig dbConfig
}
type dbConfig struct{
	conStr string
}