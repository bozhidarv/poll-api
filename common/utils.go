package common

import (
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




