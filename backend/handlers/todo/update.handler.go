package todoHandlers

import (
	"backend/db"
	"backend/db/todo"
	authHandlers "backend/handlers/auth"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

type UpdateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTodoResponse struct {
	Message string   `json:"message"`
	Todo    *db.Todo `json:"todo"`
}

// @Summary Update a todo by ID
// @Description Update a todo item for the authenticated user by its ID
// @Tags todo
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT access token"
// @Param todoId path string true "Todo ID to update"
// @Param updateReq body UpdateTodoRequest true "Updated todo details"
// @Success 200 {object} UpdateTodoResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/todo/{todoId} [put]
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	userID, ok := r.Context().Value(authHandlers.UserIdKey).(gocql.UUID)
	if !ok {
		http.Error(w, "Error getting userId", http.StatusBadRequest)
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

	var updateReq UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
