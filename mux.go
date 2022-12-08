package main

import (
	"context"
	"ent-sample/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewMux(ctx context.Context) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", HealthCheck)
	mux.Get("/todos", handlers.GetTodos)
	return mux
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, _ = w.Write([]byte(`{"status": "ok", "message": "success"}`))
}
