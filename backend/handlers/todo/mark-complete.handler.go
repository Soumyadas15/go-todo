package todoHandlers

import (
	
    "net/http"
	"backend/db"
    "encoding/json"
    "backend/db/todo"
    "github.com/gocql/gocql"
	"github.com/gorilla/mux"
)


type MarkTodoRequest struct {
    UserID string `json:"userId"`
}


type MarkTodoResponse struct {
    Message string     `json:"message"`
    Todo    *db.Todo   `json:"todo"`
}



func MarkTodoAsCompleteHandler(w http.ResponseWriter, r *http.Request) {
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

    var markCompleteReq MarkTodoRequest
    if err := json.NewDecoder(r.Body).Decode(&markCompleteReq); err != nil {
        http.Error(w, "Failed to parse request body", http.StatusBadRequest)
        return
    }

    userID, err := gocql.ParseUUID(markCompleteReq.UserID)
    if err != nil {
        http.Error(w, "Invalid userId", http.StatusBadRequest)
        return
    }

    if err := todo.MarkTodoAsComplete(todoID, userID); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := MarkTodoResponse{
        Message: "success",
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
