package polls

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/bozhidarv/poll-api/common"
	"github.com/bozhidarv/poll-api/services"
)

func GetRouter() chi.Router {
	pollsRouter := chi.NewRouter()

	pollsRouter.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		err := services.OpenDbConnection()
		if err != nil {
			common.HandleError(err, &w)
			return
		}

		defer services.CloseDbConn()

		polls, err := services.GetAllPolls()
		if err != nil {
			common.HandleError(err, &w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(polls)
		if err != nil {
			common.HandleError(err, &w)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	return pollsRouter
}
