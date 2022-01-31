package repository

import (
	"os"
	"strconv"
	"testing"
)

var (
	defaultMongoUser     = os.Getenv("MONGO_USER")
	defaultMongoPassword = os.Getenv("MONGO_PASSWORD")

	// ttNewMongoDatabase contains all tests case for NewMongoDatabase
	ttNewMongoDatabase = []struct {
		Configuration
		expectedError error
	}{
		// This case tests the connection to MongoDB using a local configuration
		{
			Configuration: Configuration{
				Host:     "localhost",
				Port:     27_017,
				Database: "test",
				User:     defaultMongoUser,
				Password: defaultMongoPassword,
			},
		},
	}
)

// TestMongoDatabase health check for MongoDB connection
func TestMongoDatabase(t *testing.T) {
	for i, v := range ttNewMongoDatabase {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			_, err := NewMongoDatabase(v.Configuration)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf(`%+v`, v.Configuration)
		})
	}
}
