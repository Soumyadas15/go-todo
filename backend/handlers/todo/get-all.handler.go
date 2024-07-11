package todoHandlers


import (
	
    "backend/db/todo"
	"net/http"
	"encoding/json"
	"log"
)



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