package health

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetRouter() chi.Router {
	testRouter := chi.NewRouter()

	testRouter.Head("", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
	})

	return testRouter
}
