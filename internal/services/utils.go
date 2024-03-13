package services

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/bozhidarv/poll-api/internal/models"
	"github.com/rs/zerolog"
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

func HandleError(err error, w *http.ResponseWriter, logger zerolog.Logger) {
	logger.Error().Msg(err.Error())
	apiError, ok := err.(*models.ApiError)
	if !ok {
		apiError = &models.ApiError{
			Code:    500,
			Message: "Internal Server Error",
		}
	}

	json.NewEncoder(*w).Encode(apiError)
	(*w).WriteHeader(apiError.Code)

	return
}
