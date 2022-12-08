package handlers

import (
	"ent-sample/repository"
	"log"
	"net/http"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todos, err := repository.DB.Todo.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}

	RespondJSON(ctx, w, todos, http.StatusOK)
}
