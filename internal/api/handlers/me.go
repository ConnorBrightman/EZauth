package handlers

import (
	"net/http"

	"ezauth/internal/httpx"
	"ezauth/internal/middleware"
)

// MeHandler returns information about the currently authenticated user
func MeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := middleware.GetUserFromContext(r)
		if !ok {
			httpx.Error(w, http.StatusUnauthorized, "unauthenticated")
			return
		}

		httpx.JSON(w, http.StatusOK, map[string]interface{}{
			"user_id": claims["user_id"],
			"email":   claims["email"],
		})
	})
}
