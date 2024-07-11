package authHandlers

import (
    "encoding/json"
    "net/http"
    "backend/db"
	"backend/db/user"
    "github.com/google/uuid"
    "github.com/gocql/gocql"
)


type RegisterResponse struct {
    Message string `json:"message"`
    User    db.User `json:"user,omitempty"`
}




func Register(w http.ResponseWriter, r *http.Request) {
    var myUser db.User
    err := json.NewDecoder(r.Body).Decode(&myUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := gocql.ParseUUID(uuid.New().String())
    if err != nil {
        http.Error(w, "failed to generate UUID", http.StatusInternalServerError)
        return
    }

    myUser.ID = id

    if err := user.CreateUser(myUser); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := RegisterResponse{
        Message: "success",
        User:    myUser,
    }
    w.Header().Set("Content-Type", "application/json");
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}
