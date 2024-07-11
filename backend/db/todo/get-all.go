package todo

import (
	"errors"
	"github.com/gocql/gocql"
    "log"
	"time"
	"backend/db"
)



func GetAllTodos() ([]db.Todo, error) {
    var todos []db.Todo

    if db.Session == nil {
        log.Println("Error fetching todos: database session not initialized")
        return todos, errors.New("database session not initialized")
    }

    iter := db.Session.Query("SELECT id, title, description, status, user_id, created_at, updated_at FROM todo_app_v1.todos").Iter()
    defer iter.Close()

    var id gocql.UUID
    var userId gocql.UUID
    var title, description string
    var status string
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
        log.Printf("Todo fetched: %+v\n", todo)
    }

    if err := iter.Close(); err != nil {
        log.Printf("Error fetching todos: %v\n", err)
        return nil, err
    }

    return todos, nil
}