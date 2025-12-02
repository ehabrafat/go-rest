package main

import (
	"myproject/internal/products"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func mount() http.Handler{
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad", 500)
	})

	r.Route("/products", func(r chi.Router){
		productHandler := products.NewHandler()
		r.Get("/", productHandler.ListProducts)
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
}
type config struct {
	addr string
}