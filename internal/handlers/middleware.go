package handlers

import (
	"context"
	"net/http"
	"strings"

	"thainsbook/internal/auth"
)

type contextKey string

const UsernameKey contextKey = "username"

func (a *Application) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			HandleUnauthorized(w, r)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			HandleUnauthorized(w, r)
			return
		}

		tokenString := parts[1]

		username, err := auth.ValidateToken(tokenString, a.JWT)
		if err != nil {
			HandleUnauthorized(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), UsernameKey, username)

		next(w, r.WithContext(ctx))
	}
}
