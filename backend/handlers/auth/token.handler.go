package authHandlers

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gocql/gocql"
)

const (
	UserIdKey = "userId"
)

func ParseToken(tokenString string) (gocql.UUID, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("saumya123456"), nil
	})
	if err != nil {
		return gocql.UUID{}, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return gocql.UUID{}, errors.New("invalid token")
	}

	return claims.UserId, nil
}

func SetUserIdInContext(r *http.Request, userId gocql.UUID) *http.Request {
	ctx := context.WithValue(r.Context(), UserIdKey, userId)
	return r.WithContext(ctx)
}
