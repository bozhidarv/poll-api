package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"github.com/bozhidarv/poll-api/common"
	"github.com/bozhidarv/poll-api/health"
	"github.com/bozhidarv/poll-api/polls"
)

func main() {
	mainRouter := chi.NewRouter()

	// Setting up middlewares
	mainRouter.Use(middleware.Logger)

	// Setting up routers
	mainRouter.Mount("/app/health", health.GetRouter())
	mainRouter.Mount("/polls", polls.GetRouter())

	// Starting the server
	PORT := common.GetPort()
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), mainRouter)
}
