package services

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/httplog"

	"github.com/bozhidarv/poll-api/internal/models"
)

func GetPort() int {
	portStr, isEnv := os.LookupEnv("PORT")

	if isEnv && portStr != "" {
		port, convErr := strconv.Atoi(portStr)
		if convErr == nil {
			return port
		}
	}
	return 3000
}

var Logger = httplog.NewLogger("poll-api")

func HandleError(err error, w *http.ResponseWriter) {
	Logger.Error().Msg(err.Error())
	apiError, ok := err.(*models.ApiError)
	if !ok {
		apiError = &models.ApiError{
			Code:    500,
			Message: "Internal Server Error",
		}
	}

	(*w).WriteHeader(apiError.Code)
	json.NewEncoder(*w).Encode(apiError)

	return
}
