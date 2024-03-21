package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bozhidarv/poll-api/internal/middleware"
	"github.com/bozhidarv/poll-api/internal/models"
	"github.com/bozhidarv/poll-api/internal/services"
	"github.com/go-chi/chi/v5"
)

func GetPollRouter() chi.Router {
	pollsRouter := chi.NewRouter()

	pollsRouter.Use(middleware.AuthMiddleware)

	pollsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("userId").(string)

		polls, err := services.GetAllPollsForUser(userId)
		if err != nil {
			services.HandleError(err, &w)
			return
		}

		err = json.NewEncoder(w).Encode(polls)
		if err != nil {
			services.HandleError(err, &w)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	pollsRouter.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		pollId := chi.URLParam(r, "id")
		poll, err := services.GetPollById(pollId)
		if err != nil {
			services.HandleError(err, &w)
			return
		}

		err = json.NewEncoder(w).Encode(poll)
		if err != nil {
			services.HandleError(err, &w)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	pollsRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		str, err := io.ReadAll(r.Body)
		if err != nil {
			services.HandleError(err, &w)
			return
		}
		var poll models.Poll
		err = json.Unmarshal(str, &poll)
		if err != nil {
			services.HandleError(err, &w)
			return
		}

		userId := r.Context().Value("userId").(string)
		poll.CreatedBy = &userId

		err = services.InsertNewPoll(poll)
		if err != nil {
			services.HandleError(err, &w)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	pollsRouter.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		pollId := chi.URLParam(r, "id")
		defer r.Body.Close()
		str, err := io.ReadAll(r.Body)
		if err != nil {
			services.HandleError(err, &w)
			return
		}
		var poll models.Poll
		err = json.Unmarshal(str, &poll)
		if err != nil {
			services.HandleError(err, &w)
			return
		}

		err = services.UpdatePoll(pollId, poll)
		if err != nil {
			services.HandleError(err, &w)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	pollsRouter.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		pollId := chi.URLParam(r, "id")

		err := services.DeletePoll(pollId)
		if err != nil {
			services.HandleError(err, &w)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	return pollsRouter
}
