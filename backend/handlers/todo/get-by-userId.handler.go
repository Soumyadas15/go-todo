package todoHandlers

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"backend/db"
	"backend/db/todo"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func GetTodoByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userId"]
	if userID == "" {
		http.Error(w, "userId parameter is required", http.StatusBadRequest)
		return
	}

	parsedUserID, err := gocql.ParseUUID(userID)
	if err != nil {
		http.Error(w, "Invalid userId", http.StatusBadRequest)
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

	todos, nextPageState, err := todo.GetTodoByUserId(parsedUserID, decodedPageState, pageSize, sortBy)
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
