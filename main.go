package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"

	"github.com/bozhidarv/poll-api/common"
	"github.com/bozhidarv/poll-api/health"
	"github.com/bozhidarv/poll-api/polls"
	"github.com/bozhidarv/poll-api/services"
)

func main() {
	mainRouter := chi.NewRouter()

	logger := httplog.NewLogger("httplog-example")

	// Setting up middlewares
	mainRouter.Use(middleware.RequestID)
	mainRouter.Use(httplog.RequestLogger(logger))
	mainRouter.Use(middleware.Recoverer)
	mainRouter.Use(middleware.AllowContentType("application/json"))

	// Setting up routers
	mainRouter.Mount("/app/health", health.GetRouter())
	mainRouter.Mount("/polls", polls.GetRouter())

	//Close db connection when the app closes
	defer services.CloseDbConn()

	// Starting the server
	PORT := common.GetPort()

	fmt.Println("Server is up and running on port:", PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mainRouter)
	if err != nil {
		fmt.Println(err.Error())
	}

}
