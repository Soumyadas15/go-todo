package todo

import (
	"backend/db"
	"log"
	"time"

	"github.com/gocql/gocql"
)

func MarkTodoAsComplete(todoId gocql.UUID, userId gocql.UUID) error {
	if db.Session == nil {
		return nil
	}

	query := db.Session.Query(`
        UPDATE todo_app_v1.todos
        SET status = ?, updated_at = ?
        WHERE id = ? AND user_id = ?
    `, "complete", time.Now(), todoId, userId)

	if err := query.Exec(); err != nil {
		log.Printf("Error updating todo with ID %s: %v", todoId.String(), err)
		return err
	}

	log.Printf("Successfully marked todo with ID %s as complete", todoId.String())
	return nil
}
