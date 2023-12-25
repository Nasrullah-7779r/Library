package database

import "go.mongodb.org/mongo-driver/mongo"

func GetDBInstance() *mongo.Database {
	return dbInstance
}
