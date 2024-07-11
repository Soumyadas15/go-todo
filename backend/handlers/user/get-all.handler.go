package userHandlers

import(

	"log" 
	"net/http"
	"encoding/json"
	"backend/db/user"
)


func GetAllUsers(w http.ResponseWriter, r *http.Request) {

    users, err := user.GetAllUsers()
    if err != nil {
        log.Printf("Failed to fetch users: %v", err)
        http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(users)
}