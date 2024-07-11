package db

import (
	"errors"
	"log"
	"time"

	"github.com/gocql/gocql"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID       gocql.UUID `json:"id" cql:"id"`
	Username string     `json:"username" cql:"username"`
	Email    string     `json:"email" cql:"email"`
	Password string     `json:"password" cql:"password"`
	Todos    []Todo     `json:"todos,omitempty"`

	CreatedAt time.Time `json:"createdAt" cql:"created_at"`
}

type Todo struct {
	ID          gocql.UUID `json:"id" cql:"id"`
	Title       string     `json:"title" cql:"title"`
	Description string     `json:"description" cql:"description"`
	Status      string     `json:"status" cql:"status"`
	UserID      gocql.UUID `json:"userId" cql:"user_id"`

	CreatedAt time.Time `json:"createdAt" cql:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" cql:"updated_at"`
}

var Session *gocql.Session

func InitCluster(databaseURI string) {

	var err error

	var cluster = gocql.NewCluster(databaseURI)

	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	if err := createKeyspaceAndTables(Session); err != nil {
		log.Fatalf("Failed to create keyspace and tables: %v", err)
	}

	log.Println("Db connected")

}

func createKeyspaceAndTables(session *gocql.Session) error {

	if err := session.Query(`
        CREATE KEYSPACE IF NOT EXISTS todo_app_v1
        WITH replication = {
          'class': 'SimpleStrategy',
          'replication_factor': 3
        };
    `).Exec(); err != nil {
		return err
	}

	if err := Session.Query(`
        CREATE TABLE IF NOT EXISTS todo_app_v1.todos (
            user_id UUID,
            created_at TIMESTAMP,
            id UUID,
            title TEXT,
            description TEXT,
            status TEXT,
            updated_at TIMESTAMP,
            PRIMARY KEY (user_id, id)
        ) WITH CLUSTERING ORDER BY (id ASC);
    `).Exec(); err != nil {
		return err
	}

	if err := session.Query(`
        CREATE TABLE IF NOT EXISTS todo_app_v1.users (
            id UUID PRIMARY KEY,
            username TEXT,
            email TEXT,
            password TEXT,
            todos LIST<UUID>,
            created_at TIMESTAMP
        );
    `).Exec(); err != nil {
		return err
	}

	return nil
}

func CloseCluster() {
	if Session != nil {
		Session.Close()
		log.Println("Database session closed")
	}
}
