// Package repository contains everything related to storage
package repository

import "fmt"

// Type defines which repository be used
type Type uint

// Defines correct values for Type
const (
	// NoSQL defines an NoSQL storage oriented to documents like MongoDB
	NoSQL Type = iota
)

// _ implements constraint for Configuration struct
var _ fmt.Stringer = Configuration{}

// Configuration settings used to make a repository connection
type Configuration struct {
	// Type defines the type of repository to which you want to establish a connection
	Type
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

// String build and returns a URI used to establish a connection to any repository defined by the Type embbed in the Configuration structure
//
// Supported types: NoSQL
func (c Configuration) String() string {
	switch c.Type {
	case NoSQL:
		return fmt.Sprintf(
			"mongodb://%s:%s@%s:%d", //?maxPoolSize=%s",
			c.User,
			c.Password,
			c.Host,
			c.Port,
			// os.Getenv("MONGO_POOL_SIZE"),
		)
	}

	panic(fmt.Sprintf(`type "%d" is not supported by Configuration.String()`, c.Type))
}
