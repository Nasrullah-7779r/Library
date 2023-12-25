package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Client
var dbInstance *mongo.Database

//var dbInstance DB
