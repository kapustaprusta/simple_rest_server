package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// TestMain ...
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "user=vladislav password=24021996 dbname=restapi_dev sslmode=disable"
	}

	os.Exit(m.Run())
}
