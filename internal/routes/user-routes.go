package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bozhidarv/poll-api/internal/models"
	"github.com/bozhidarv/poll-api/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
)

func GetUserRouter() chi.Router {
	userRoutes := chi.NewRouter()

	userRoutes.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		logger := httplog.LogEntry(r.Context())
		defer r.Body.Close()
		str, err := io.ReadAll(r.Body)
		if err != nil {
			services.HandleError(err, &w, logger)
			return
		}
		var user models.User
		err = json.Unmarshal(str, &user)
		if err != nil {
			services.HandleError(err, &w, logger)
			return
		}

		err = services.RegisterUser(user)
		if err != nil {
			services.HandleError(err, &w, logger)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	return userRoutes
}
