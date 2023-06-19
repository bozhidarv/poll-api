package common

import (
	"net/http"
	"os"
	"strconv"
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

func HandleError(err error, w *http.ResponseWriter) {
	(*w).WriteHeader(500)
	(*w).Write([]byte(err.Error()))
	return
}
