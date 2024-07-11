package authHandlers

import (
	"backend/db"
	"backend/db/user"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterResponse struct {
	Message string  `json:"message"`
	User    db.User `json:"user,omitempty"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary Register
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User object to register"
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/register [post]
func Register(w http.ResponseWriter, r *http.Request) {

	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := gocql.ParseUUID(uuid.New().String())
	if err != nil {
		http.Error(w, "failed to generate UUID", http.StatusInternalServerError)
		return
	}

	req.Email = strings.TrimSpace(req.Email)
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)

	if req.Email == "" || req.Username == "" || req.Password == "" {
		http.Error(w, "Fields cannot be null", http.StatusBadRequest)
		return
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return
	}

	myUser := db.User{
		Email:     req.Email,
		Username:  req.Username,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
