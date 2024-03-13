package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetRouter() chi.Router {
	healthRouter := chi.NewRouter()

	healthRouter.Head("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
	})

	return healthRouter
}
