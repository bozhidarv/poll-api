package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/bozhidarv/poll-api/internal/models"
	"github.com/bozhidarv/poll-api/internal/services"
)

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		unauthorizedError := &models.ApiError{
			Code:    401,
			Message: "Unauthorized",
		}

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			services.HandleError(unauthorizedError, &w)
			return
		}

		token := strings.Split(authHeader, " ")[1]

		userId, expTimestamp := services.ParseJwtToken(token)

		if expTimestamp < float64(time.Now().Unix()) {
			services.HandleError(unauthorizedError, &w)
			return
		}

		newCtx := context.WithValue(r.Context(), "userId", userId)

		handler.ServeHTTP(w, r.WithContext(newCtx))
	})
}
