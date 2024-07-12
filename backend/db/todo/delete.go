package todo

import (
	"backend/db"
	"log"

	"github.com/gocql/gocql"
)

func DeleteTodoByID(todoID gocql.UUID, userId gocql.UUID) error {
	if db.Session == nil {
		return nil
	}

	query := db.Session.Query("DELETE FROM todo_app_v1.todos WHERE id = ? AND user_id = ?", todoID, userId)

	if err := query.Exec(); err != nil {
		log.Printf("Error deleting todo with ID %s: %v", todoID.String(), err)
		return err
	}

	log.Printf("Todo with ID %s deleted successfully", todoID.String())

	return nil
}
