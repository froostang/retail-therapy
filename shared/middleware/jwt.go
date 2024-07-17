package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

// TODO: add logging component?
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractTokenFromHeader(r)
		if tokenString == "" {
			// http.Error(w, "Unauthorized", http.StatusUnauthorized)
			http.Redirect(w, r, "/unauthorized", http.StatusFound)
			return
		}

		// TODO: redirect on bad auth correctly for login
		token, err := validateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
	})
}

func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return ""
	}

	return tokenParts[1]
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the token signing method here
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
