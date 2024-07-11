package todo

import (
	"github.com/gocql/gocql"
    "log"
	"backend/db"
)

func UpdateTodoByID(todo db.Todo, userId gocql.UUID) error {
    if db.Session == nil {
        return nil
    }

    query := db.Session.Query(`
        UPDATE todo_app_v1.todos
        SET title = ?, description = ?, updated_at = ?
        WHERE id = ? AND user_id = ?
    `, todo.Title, todo.Description, todo.UpdatedAt, todo.ID, userId)

    if err := query.Exec(); err != nil {
        log.Printf("Error updating todo with ID %s: %v", todo.ID.String(), err)
        return err
    }

    log.Printf("Todo with ID %s updated successfully", todo.ID.String())
    return nil
}