package user

import (
	"backend/db"
	"errors"
)

func CreateUser(user db.User) error {
	if db.Session == nil {
		return errors.New("database session not initialized")
	}

	_, err := GetUserByEmail(user.Email)
	if err == nil {
		return errors.New("user with this email already exists")
	} else if err != nil && err.Error() != "user not found" {
		return err
	}

	_, err = GetUserByUsername(user.Username)
	if err == nil {
		return errors.New("user with this username already exists")
	} else if err != nil && err.Error() != "user not found" {
		return err
	}

	query := db.Session.Query(`
        INSERT INTO todo_app_v1.users (id, username, email, password, todos, created_at)
        VALUES (?, ?, ?, ?, ?, ?)
    `, user.ID, user.Username, user.Email, user.Password, user.Todos, user.CreatedAt)

	if err := query.Exec(); err != nil {
		return err
	}

	return nil
}
