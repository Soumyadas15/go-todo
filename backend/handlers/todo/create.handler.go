package todoHandlers


import (
	"time"
	"backend/db"
    "backend/db/todo"
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
    "github.com/gocql/gocql"
)


type CreateTodoResponse struct {
    Message string    `json:"message"`
    Todo    db.Todo   `json:"todo"`
}


func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
    var myTodo db.Todo

    err := json.NewDecoder(r.Body).Decode(&myTodo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	currentTime := time.Now()
    myTodo.CreatedAt = currentTime
    myTodo.UpdatedAt = currentTime
	myTodo.Status = "pending"

	id, err := gocql.ParseUUID(uuid.New().String())
	if err != nil {
        http.Error(w, "failed to generate UUID", http.StatusInternalServerError)
        return
    }

	myTodo.ID = id

    if err := todo.CreateTodo(myTodo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := CreateTodoResponse{
        Message: "success",
        Todo:    myTodo,
    }

    w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}