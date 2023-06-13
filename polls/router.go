package polls

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/bozhidarv/poll-api/services"
)

func GetRouter() chi.Router {
	pollsRouter := chi.NewRouter()

	pollsRouter.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		services.OpenDbConnection()
		defer services.CloseDbConn()
		polls, err := services.GetAllPolls()
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}
		fmt.Println(len(polls))
		json.NewEncoder(w).Encode(polls)
		w.WriteHeader(200)
	})

	return pollsRouter
}
