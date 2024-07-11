package todoHandlers

import (
	"backend/db"
	"backend/db/todo"
	authHandlers "backend/handlers/auth"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTodoResponse struct {
	Message string  `json:"message"`
	Todo    db.Todo `json:"todo"`
}

// @Summary Create a new todo
// @Description Create a new todo item for the authenticated user
// @Tags todo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT access token"
// @Param todo body CreateTodoRequest true "Todo object to be created"
// @Success 201 {object} CreateTodoResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/todo [post]
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(authHandlers.UserIdKey).(gocql.UUID)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var reqBody CreateTodoRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reqBody.Title = strings.TrimSpace(reqBody.Title)
	reqBody.Description = strings.TrimSpace(reqBody.Description)

	if reqBody.Title == "" || reqBody.Description == "" {
		http.Error(w, "Title and Description cannot be empty", http.StatusBadRequest)
		return
	}

	myTodo := db.Todo{
		Title:       reqBody.Title,
		Description: reqBody.Description,
		UserID:      userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      "pending",
		ID:          gocql.TimeUUID(),
	}

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
