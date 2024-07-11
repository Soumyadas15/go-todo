package todoHandlers

import (
	"backend/db"
	"backend/db/todo"
	"encoding/json"
	"log"
	"net/http"
)

type Todo struct {
	db.Todo
}

// @Summary Get all todo items
// @Description Retrieve all todo items from the database
// @Tags todo
// @Produce  json
// @Success 200 {array} db.Todo
// @Failure 500 {object} map[string]string
// @Router /api/todos [get]
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := todo.GetAllTodos()
	if err != nil {
		log.Printf("Failed to fetch todos: %v", err)
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
