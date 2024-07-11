package todoHandlers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"backend/db"
	"backend/db/todo"
	authHandlers "backend/handlers/auth"

	"github.com/gocql/gocql"
)

type GetTodosResponse struct {
	Todos         []db.Todo `json:"todos"`
	NextPageState string    `json:"nextPageState,omitempty"`
}

// @Summary Get todo items by user ID
// @Description Retrieve todo items belonging to the authenticated user
// @Tags todo
// @Produce  json
// @Param Authorization header string true "JWT access token"
// @Param pageSize query integer false "Number of items per page (default is 4)"
// @Param pageState query string false "Page state for pagination"
// @Param sortBy query string false "Field to sort by"
// @Success 200 {object} GetTodosResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/user/todos [get]
func GetTodoByUserId(w http.ResponseWriter, r *http.Request) {

	var err error

	userID, ok := r.Context().Value(authHandlers.UserIdKey).(gocql.UUID)
	if !ok {
		http.Error(w, "Unable to parse userId", http.StatusBadRequest)
		return
	}

	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize := 0

	pageState := r.URL.Query().Get("pageState")

	if pageSizeStr != "" {
		pageSize, err = strconv.Atoi(pageSizeStr)
		if err != nil {
			http.Error(w, "Invalid pageSize", http.StatusBadRequest)
			return
		}
		if pageSize > 4 {
			pageSize = 4
		}
	} else {
		pageSize = 4
	}

	var decodedPageState []byte
	if pageState != "" {
		decodedPageState, err = base64.RawURLEncoding.DecodeString(pageState)
		if err != nil {
			http.Error(w, "Invalid pageState", http.StatusBadRequest)
			return
		}
	}

	sortBy := r.URL.Query().Get("sortBy")

	todos, nextPageState, err := todo.GetTodoByUserId(userID, decodedPageState, pageSize, sortBy)
	if err != nil {
		log.Printf("Failed to fetch todos: %v", err)
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var encodedNextPageState string
	if nextPageState != nil {
		encodedNextPageState = base64.RawURLEncoding.EncodeToString(nextPageState)
	}

	response := struct {
		Todos         []db.Todo `json:"todos"`
		NextPageState string    `json:"nextPageState,omitempty"`
	}{
		Todos:         todos,
		NextPageState: encodedNextPageState,
	}

	json.NewEncoder(w).Encode(response)
}
