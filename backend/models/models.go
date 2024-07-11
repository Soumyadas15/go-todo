package models

import (
    "time"
    "github.com/gocql/gocql"
)

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
    CreatedAt   time.Time  `json:"createdAt" cql:"created_at"`
    UpdatedAt   time.Time  `json:"updatedAt" cql:"updated_at"`
}
