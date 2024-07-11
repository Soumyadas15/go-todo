package todoHandlers

import (
	"backend/db/todo"
	authHandlers "backend/handlers/auth"
	"encoding/json"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

type DeleteResponse struct {
	Message string `json:"message"`
}

// @Summary Delete a todo by ID
// @Description Delete a todo item belonging to the authenticated user by its ID
// @Tags todo
// @Produce  json
// @Param Authorization header string true "JWT access token"
// @Param todoId path string true "Todo ID to delete"
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/todo/{todoId} [delete]
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	todoIDStr, ok := vars["todoId"]
	if !ok {
		http.Error(w, "Missing todoId in URL", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(authHandlers.UserIdKey).(gocql.UUID)
	if !ok {
		http.Error(w, "Unable to parse userId", http.StatusBadRequest)
		return
	}

	todoID, err := gocql.ParseUUID(todoIDStr)
	if err != nil {
		http.Error(w, "Invalid todoId", http.StatusBadRequest)
		return
	}

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
