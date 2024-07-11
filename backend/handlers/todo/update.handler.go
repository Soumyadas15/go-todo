package todoHandlers

import (
	"time"
    "net/http"
	"backend/db"
    "encoding/json"
    "backend/db/todo"
    "github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

type UpdateTodoRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    UserID      string `json:"userId"`
}


type UpdateTodoResponse struct {
    Message string     `json:"message"`
    Todo    *db.Todo   `json:"todo"`
}



func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    todoIDStr, ok := vars["todoId"]
    if !ok {
        http.Error(w, "Missing todoId in URL", http.StatusBadRequest)
        return
    }

    todoID, err := gocql.ParseUUID(todoIDStr)
    if err != nil {
        http.Error(w, "Invalid todoId", http.StatusBadRequest)
        return
    }

    var updateReq UpdateTodoRequest
    if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    userID, err := gocql.ParseUUID(updateReq.UserID)
    if err != nil {
        http.Error(w, "Invalid userId", http.StatusBadRequest)
        return
    }

    existingTodo, err := todo.GetTodoByID(todoID, userID)
    if err != nil {
        http.Error(w, "Todo not found", http.StatusNotFound)
        return
    }

    existingTodo.Title = updateReq.Title
    existingTodo.Description = updateReq.Description
    existingTodo.UpdatedAt = time.Now()

    if err := todo.UpdateTodoByID(*existingTodo, userID); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := UpdateTodoResponse{
        Message: "success",
        Todo:    existingTodo,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}