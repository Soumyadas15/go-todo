package todo

import (
	"backend/db"
	"errors"
	"log"

	"github.com/gocql/gocql"
)

func GetTodoByID(todoID gocql.UUID, userID gocql.UUID) (*db.Todo, error) {
	if db.Session == nil {
		return nil, errors.New("database session not initialized")
	}

	var todo db.Todo
	query := db.Session.Query(`
        SELECT id, title, description, status, user_id, created_at, updated_at
        FROM todo_app_v1.todos
        WHERE id = ? AND user_id = ?
    `, todoID, userID)

	if err := query.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.UserID, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		log.Printf("Error fetching todo with ID %s: %v", todoID.String(), err)
		return nil, err
	}

	return &todo, nil
}
