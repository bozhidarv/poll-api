package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bozhidarv/poll-api/internal/middleware"
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

		if _, err = services.GetUserByEmail(*user.Email); err == nil {
			apiErr := &models.ApiError{
				Code:    400,
				Message: "User with this email already exists.",
			}
			services.HandleError(apiErr, &w)
			return
		}

		userId, err := services.CreateUser(user)
		if err != nil {
			services.HandleError(err, &w)
			return
		}

		token := services.CreateJwtToken(userId)

		w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))

		w.WriteHeader(http.StatusOK)
	})

	userRoutes.Post("/login", func(w http.ResponseWriter, r *http.Request) {
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

		dbUser, err := services.GetUserByEmail(*user.Email)
		if err != nil {
			services.HandleError(err, &w)
			return
		}

		if !services.CheckPasswordValidity(*user.Password, *dbUser.Password) {
			apiError := &models.ApiError{
				Code:    400,
				Message: "Wrong password.",
			}

			services.HandleError(apiError, &w)
			return
		}

		token := services.CreateJwtToken(*dbUser.Id)

		w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))

		w.WriteHeader(http.StatusOK)
	})

	return userRoutes
}

func (*UserRoutes) GetProtectedUserRoutes() chi.Router {
	userRoutes := chi.NewRouter()
	userRoutes.Use(middleware.AuthMiddleware)

	userRoutes.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Authorization", "")
		w.WriteHeader(http.StatusOK)
	})

	return userRoutes
}
