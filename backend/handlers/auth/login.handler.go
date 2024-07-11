package authHandlers

import (
    "encoding/json"
    "net/http"
    "backend/db"
	"backend/db/user"
    "log"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)


type Response struct {
    Message string `json:"message"`
    User    db.User `json:"user,omitempty"`
    Token   string      `json:"token,omitempty"`
}

type Credentials struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}



func Login(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    log.Printf("Querying user with email: %s\n", creds.Email)
    log.Printf("Querying user with email: %s\n", creds.Password)

    user, err := user.GetUserByEmail(creds.Email)
    if err != nil {
        if err != db.ErrUserNotFound {
            http.Error(w, "User does not exist", http.StatusUnauthorized)
        } else {
            http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
        }
        return
    }

    if !comparePasswords(user.Password, creds.Password) {
        http.Error(w, "Incorrect password", http.StatusUnauthorized)
        return
    }

	token, err := createToken(user.Email)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    response := Response{
        Message: "success",
        User:    db.User{ID: user.ID, Email: user.Email, Username: user.Username},
		Token:   token,
    }
    
    w.Header().Set("Content-Type", "application/json");
    w.Header().Set("Access-Control-Allow-Origin", "*")
    json.NewEncoder(w).Encode(response)
}


func createToken(email string) (string, error) {
    claims := Claims{
        Email: email,
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


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func VerifyPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


func comparePasswords(savedPassword, inputPassword string) bool {
    return savedPassword == inputPassword
}