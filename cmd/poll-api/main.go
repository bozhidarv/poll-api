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

	mainRouter.Use(middleware.RequestID)
	mainRouter.Use(httplog.RequestLogger(services.Logger))
	mainRouter.Use(middleware.Recoverer)
	mainRouter.Use(middleware.AllowContentType("application/json"))
	mainRouter.Use(middleware.SetHeader("Content-Type", "application/json"))

	mainRouter.Mount("/polls", routes.GetPollRouter())
	userRouter := routes.UserRoutes{}
	mainRouter.Mount("/", userRouter.GetUnprotectedUserRouter())

	defer func() {
		err := services.CloseDbConn()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	PORT := services.GetPort()

	fmt.Println("Server is up and running on port:", PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mainRouter)
	if err != nil {
		fmt.Println(err.Error())
	}
}
