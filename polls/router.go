package polls

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"

	"github.com/bozhidarv/poll-api/common"
	"github.com/bozhidarv/poll-api/polls/models"
	"github.com/bozhidarv/poll-api/services"
)

func GetRouter() chi.Router {
	pollsRouter := chi.NewRouter()

	pollsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		logger := httplog.LogEntry(r.Context())
		polls, err := services.GetAllPolls()
		if err != nil {
			common.HandleError(err, &w, logger)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(polls)
		if err != nil {
			common.HandleError(err, &w, logger)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	pollsRouter.Post("/", func(w http.ResponseWriter, r *http.Request) {
		logger := httplog.LogEntry(r.Context())
		defer r.Body.Close()
		str, err := io.ReadAll(r.Body)
		if err != nil {
			common.HandleError(err, &w, logger)
			return
		}
		var poll models.Poll
		err = json.Unmarshal(str, &poll)
		if err != nil {
			common.HandleError(err, &w, logger)
			return
		}

		services.InsertNewPoll(poll)
		if err != nil {
			common.HandleError(err, &w, logger)
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
			common.HandleError(err, &w, logger)
			return
		}
		var poll models.Poll
		err = json.Unmarshal(str, &poll)
		if err != nil {
			common.HandleError(err, &w, logger)
			return
		}

		services.UpdatePoll(pollId, poll)
		if err != nil {
			common.HandleError(err, &w, logger)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	return pollsRouter
}
