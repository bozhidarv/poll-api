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

func GetPollRouter() chi.Router {
	pollsRouter := chi.NewRouter()

	pollsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		logger := httplog.LogEntry(r.Context())
		polls, err := services.GetAllPolls()
		if err != nil {
			services.HandleError(err, &w, logger)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(polls)
		if err != nil {
			services.HandleError(err, &w, logger)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	pollsRouter.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		pollId := chi.URLParam(r, "id")
		logger := httplog.LogEntry(r.Context())
		poll, err := services.GetPollById(pollId)
		if err != nil {
			if err.Error() == "NOT_FOUND" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			services.HandleError(err, &w, logger)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(poll)
		if err != nil {
			services.HandleError(err, &w, logger)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	pollsRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		logger := httplog.LogEntry(r.Context())
		defer r.Body.Close()
		str, err := io.ReadAll(r.Body)
		if err != nil {
			services.HandleError(err, &w, logger)
			return
		}
		var poll models.Poll
		err = json.Unmarshal(str, &poll)
		if err != nil {
			services.HandleError(err, &w, logger)
			return
		}

		err = services.InsertNewPoll(poll)
		if err != nil {
			services.HandleError(err, &w, logger)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	pollsRouter.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		pollId := chi.URLParam(r, "id")
		logger := httplog.LogEntry(r.Context())
		defer r.Body.Close()
		str, err := io.ReadAll(r.Body)
		if err != nil {
			services.HandleError(err, &w, logger)
			return
		}
		var poll models.Poll
		err = json.Unmarshal(str, &poll)
		if err != nil {
			services.HandleError(err, &w, logger)
			return
		}

		err = services.UpdatePoll(pollId, poll)
		if err != nil {
			if err.Error() == "NOT_FOUND" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			services.HandleError(err, &w, logger)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	pollsRouter.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		pollId := chi.URLParam(r, "id")
		logger := httplog.LogEntry(r.Context())

		err := services.DeletePoll(pollId)
		if err != nil {
			if err.Error() == "NOT_FOUND" {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			services.HandleError(err, &w, logger)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	return pollsRouter
}
