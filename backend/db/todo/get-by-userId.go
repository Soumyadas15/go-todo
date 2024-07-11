package todo

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"backend/db"

	"github.com/gocql/gocql"
)

func GetTodoByUserId(userId gocql.UUID, pageState []byte, pageSize int, sortBy string, sortByTime bool) ([]db.Todo, []byte, error) {
	var todos []db.Todo
	var nextPageState []byte

	if db.Session == nil {
		return todos, nextPageState, errors.New("database session not initialized")
	}

	query := "SELECT id, title, description, status, user_id, created_at, updated_at FROM todo_app_v1.todos WHERE user_id = ?"
	var args []interface{}
	args = append(args, userId)

	if sortBy != "" {
		query += " AND status = ?"
		args = append(args, sortBy)
	}
	query += " ALLOW FILTERING"

	iter := db.Session.Query(query, args...).PageSize(pageSize).PageState(pageState).Iter()
	defer iter.Close()

	var id gocql.UUID
	var title, description, status string
	var createdAt, updatedAt time.Time

	for iter.Scan(&id, &title, &description, &status, &userId, &createdAt, &updatedAt) {
		todo := db.Todo{
			ID:          id,
			Title:       title,
			Description: description,
			Status:      status,
			UserID:      userId,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}
		todos = append(todos, todo)
	}

	nextPageState = iter.PageState()

	if err := iter.Close(); err != nil {
		return nil, nextPageState, fmt.Errorf("error fetching todos: %v", err)
	}

	if sortByTime {
		sort.Slice(todos, func(i, j int) bool {
			return todos[i].CreatedAt.Before(todos[j].CreatedAt)
		})
	}

	return todos, nextPageState, nil
}
