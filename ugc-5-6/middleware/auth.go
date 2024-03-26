package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type CustomClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		claims := &CustomClaims{}
		tkn, err := jwt.ParseWithClaims(tokenParts[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret-key"), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Invalid token signature", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if !tkn.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		r = r.WithContext(NewContextWithRole(r.Context(), claims.Role))

		next(w, r, p)
	}
}

func AdminOnlyMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		role := GetRoleFromContext(r.Context())
		if role != "superAdmin" {
			http.Error(w, "Access denied", http.StatusForbidden)
			return
		}
		next(w, r, p)
	}
}

func NewContextWithRole(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, "role", role)
}

func GetRoleFromContext(ctx context.Context) string {
	role, _ := ctx.Value("role").(string)
	return role
}
