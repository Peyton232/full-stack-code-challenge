package auth

import (
	"net/http"

	"github.com/Peyton232/scoir-server/util"
)

func (svc *AuthService) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isValidToken(r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}

// isValidToken checks if the provided authentication token is not empty
func isValidToken(r *http.Request) bool {
	token := util.GetTokenFromRequest(r)
	return token != ""
}
