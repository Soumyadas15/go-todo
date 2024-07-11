package user

import (

	"backend/db"
	"errors"
    "log"
)

func GetUserByUsername(username string) (db.User, error) {
    var user db.User

    if db.Session == nil {
        return user, errors.New("database session not initialized")
    }

    query := db.Session.Query(`
        SELECT id, username, email, password, todos, created_at
        FROM todo_app_v1.users
        WHERE username = ?
        ALLOW FILTERING
    `, username)

    iter := query.Iter()
    defer iter.Close()

    if !iter.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Todos, &user.CreatedAt) {
        if err := iter.Close(); err != nil {
            log.Printf("Error fetching user: %v\n", err)
            return user, err
        }
        return user, errors.New("user not found")
    }

    return user, nil
}