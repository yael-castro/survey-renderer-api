package repository

import (
	"strconv"
	"testing"
)

var (
	defaultSettings = Configuration{
		Host:     "test.efago.mongodb.net",
		Database: "test",
		User:     "tester",
		Password: "tester",
		Secure:   true, // +srv = port 27_018
	}

	localSettings = Configuration{
		Host:     "localhost",
		Port:     27_017, // default port "27_017"
		Database: "test",
		User:     "tester",
		Password: "tester",
	}

	// ttNewMongoDatabase contains all tests case for NewMongoDatabase
	ttNewMongoDatabase = []struct {
		Configuration
		expectedError error
	}{
		// This case tests the remote connection to MongoDB
		{Configuration: defaultSettings},
		// This case tests the connection to MongoDB using a local configuration
		{Configuration: localSettings},
	}
)

// TestNewMongoDatabase health check for MongoDB connections
func TestNewMongoDatabase(t *testing.T) {
	for i, v := range ttNewMongoDatabase {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			_, err := NewMongoDatabase(v.Configuration)
			if err != nil {
				t.Fatalf(`%v: %+v`, err, v.Configuration)
			}

			t.Logf(`%+v`, v.Configuration)
		})
	}
}
