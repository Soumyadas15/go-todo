package todo

import (
    "backend/db"
	"errors"
)


func CreateTodo(todo db.Todo) error {
    if db.Session == nil {
        return errors.New("database session not initialized")
    }

    query := db.Session.Query(`
        INSERT INTO todo_app_v1.todos (id, title, description, status, user_id, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `, todo.ID, todo.Title, todo.Description, todo.Status, todo.UserID, todo.CreatedAt, todo.UpdatedAt)

    if err := query.Exec(); err != nil {
        return err
    }

    return nil
}