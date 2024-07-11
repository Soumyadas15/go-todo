package authHandlers

import (
	"backend/db"
	"backend/db/user"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Message string  `json:"message"`
	User    db.User `json:"user,omitempty"`
	Token   string  `json:"token,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	UserId gocql.UUID `json:"user_id"`
	Email  string     `json:"email"`
	jwt.StandardClaims
}

// @Summary Login
// @Description Authenticate user credentials and generate JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param creds body LoginRequest true "User credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/auth/login [post]
func Login(w http.ResponseWriter, r *http.Request) {

	var creds LoginRequest
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Querying user with email: %s\n", creds.Email)

	user, err := user.GetUserByEmail(creds.Email)
	if err != nil {
		if err != db.ErrUserNotFound {
			http.Error(w, "User does not exist", http.StatusUnauthorized)
		} else {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		}
		return
	}

	if !CheckPasswordHash(creds.Password, user.Password) {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	token, err := createToken(user.ID, user.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		Message: "success",
		User:    db.User{ID: user.ID, Email: user.Email, Username: user.Username},
		Token:   token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func createToken(userId gocql.UUID, email string) (string, error) {
	claims := Claims{
		UserId: userId,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte("saumya123456")
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
