package middleware

import (
	authHandlers "backend/handlers/auth"
	"net/http"
	"strings"
)

func VerifyToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing auth token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		userId, err := authHandlers.ParseToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		r = authHandlers.SetUserIdInContext(r, userId)

		next.ServeHTTP(w, r)
	}
}
