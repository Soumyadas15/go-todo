package todoHandlers

import (
	"backend/db"
	"backend/db/todo"
	authHandlers "backend/handlers/auth"
	"encoding/json"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

type MarkTodoResponse struct {
	Message string   `json:"message"`
	Todo    *db.Todo `json:"todo"`
}

// @Summary Mark a todo as complete
// @Description Mark a todo item as complete for the authenticated user by its ID
// @Tags todo
// @Produce  json
// @Param Authorization header string true "JWT access token"
// @Param todoId path string true "Todo ID to mark as complete"
// @Success 200 {object} MarkTodoResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/todo/{todoId}/complete [put]
func MarkTodoAsCompleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, ok := r.Context().Value(authHandlers.UserIdKey).(gocql.UUID)
	if !ok {
		http.Error(w, "Missing todoId in URL", http.StatusBadRequest)
		return
	}

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
