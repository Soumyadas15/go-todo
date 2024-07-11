package user

import (
    "errors"
	"github.com/gocql/gocql"
    "log"
	"time"
	"backend/db"
)




func GetAllUsers() ([]db.User, error) {
    var users []db.User

    if db.Session == nil {
        log.Println("Error fetching users: database session not initialized")
        return users, errors.New("database session not initialized")
    }

    iter := db.Session.Query("SELECT id, username, email, password, todos, created_at FROM todo_app_v1.users").Iter()
    defer iter.Close()

    var id gocql.UUID
    var username, email, password string
    var todos []db.Todo
    var createdAt time.Time

    for iter.Scan(&id, &username, &email, &password, &todos, &createdAt) {
        user := db.User{
            ID:         id,
            Username:   username,
            Email:      email,
            Password:   password,
            Todos:      todos,
            CreatedAt:  createdAt,
        }
        users = append(users, user)
        log.Printf("User fetched: %+v\n", user)
    }

    if err := iter.Close(); err != nil {
        log.Printf("Error fetching users: %v\n", err)
        return nil, err
    }

    return users, nil
}