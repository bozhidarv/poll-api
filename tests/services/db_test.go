package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bozhidarv/poll-api/services"
)

func testDbConnection(t *testing.T) {
	assert := assert.New(t)
	err := services.OpenDbConnection()
	defer services.CloseDbConn()

	assert.Equal(err, nil)
}
