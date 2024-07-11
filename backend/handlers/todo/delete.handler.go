package todoHandlers

import (
    "net/http"
    "encoding/json"
    "backend/db/todo"
    "github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

type DeleteRequest struct {
    UserID    string   `json:"userId"`
}

type DeleteResponse struct {
    Message   string  `json:"message"`
}


func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
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

    var deleteReq DeleteRequest
    if err := json.NewDecoder(r.Body).Decode(&deleteReq); err != nil {
        http.Error(w, "Failed to parse request body", http.StatusBadRequest)
        return
    }

    userID, err := gocql.ParseUUID(deleteReq.UserID)
    if err != nil {
        http.Error(w, "Invalid userId", http.StatusBadRequest)
        return
    }

    if err := todo.DeleteTodoByID(todoID, userID); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := DeleteResponse{Message: "success"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}