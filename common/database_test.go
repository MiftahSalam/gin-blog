package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {
	asserts := assert.New(t)
	db := Init()

	sqlDB, err := db.DB()

	asserts.NoError(err, "Db should be available")
	asserts.NoError(sqlDB.Ping(), "Db should be connected")
}
