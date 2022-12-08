package handlers

import (
	"ent-sample/repository"
	"fmt"
	"log"
	"net/http"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	items, err := repository.DB.Todo.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}

	for _, t := range items {
		fmt.Printf("%d: %q %v\n", t.ID, t.Text, t.CreatedAt)
	}

	RespondJSON(ctx, w, items, http.StatusOK)
}
