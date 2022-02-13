package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewMongoClient establish connection with a MongoDB database using the Configuration passed as parameter
func NewMongoDatabase(config Configuration) (mongoDatabase *mongo.Database, err error) {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.String()))
	if err != nil {
		return
	}

	err = mongoClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return
	}

	mongoDatabase = mongoClient.Database(config.Database)
	err = mongoDatabase.Client().Ping(context.TODO(), readpref.Primary())

	return
}

// NewMD works same that function NewMongoDatabase but it does not return an error instead it panics
func NewMD(config Configuration) *mongo.Database {
	mongoDatabase, err := NewMongoDatabase(config)
	if err != nil {
		panic(err)
	}

	return mongoDatabase
}
