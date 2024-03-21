package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bozhidarv/poll-api/internal/models"
	"github.com/bozhidarv/poll-api/internal/services"
	"github.com/go-chi/chi/v5"
)

type UserRoutes struct{}

func (*UserRoutes) GetUnprotectedUserRouter() chi.Router {
	userRoutes := chi.NewRouter()

	userRoutes.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		str, err := io.ReadAll(r.Body)
		if err != nil {
			services.HandleError(err, &w)
			return
		}
		var user models.User
		err = json.Unmarshal(str, &user)
		if err != nil {
			services.HandleError(err, &w)
			return
		}

		userId, err := services.RegisterUser(user)
		if err != nil {
			services.HandleError(err, &w)
			return
		}

		token := services.CreateJwtToken(userId)

		w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))

		w.WriteHeader(http.StatusOK)
	})

	return userRoutes
}

func (*UserRoutes) GetProtectedUserRoutes() chi.Router {
	userRoutes := chi.NewRouter()

	return userRoutes
}
