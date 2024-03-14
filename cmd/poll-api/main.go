package main

import (
	"fmt"
	"net/http"

	"github.com/bozhidarv/poll-api/internal/routes"
	"github.com/bozhidarv/poll-api/internal/services"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
)

func main() {
	mainRouter := chi.NewRouter()

	logger := httplog.NewLogger("poll-api")

	// Setting up middlewares
	mainRouter.Use(middleware.RequestID)
	mainRouter.Use(httplog.RequestLogger(logger))
	mainRouter.Use(middleware.Recoverer)
	mainRouter.Use(middleware.AllowContentType("application/json"))
	mainRouter.Use(middleware.SetHeader("Content-Type", "application/json"))

	// Setting up routers
	mainRouter.Mount("/app/health", routes.GetRouter())
	mainRouter.Mount("/polls", routes.GetPollRouter())
	mainRouter.Mount("/users", routes.GetUserRouter())

	// Close db connection when the app closes
	defer func() {
		err := services.CloseDbConn()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	// Starting the server
	PORT := services.GetPort()

	fmt.Println("Server is up and running on port:", PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mainRouter)
	if err != nil {
		fmt.Println(err.Error())
	}
}
