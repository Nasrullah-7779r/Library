package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

//var DB *mongo.Database

func Connect() {

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_CONN_STRING"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	// Set up the default database
	dbInstance = client.Database("library")

	//fmt.Println("host is", db.Host)
	//connConfig := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", db.Host, db.User, db.Password, db.Dbname, db.Port)
	//
	//dbSession, err := gorm.Open(postgres.Open(connConfig), &gorm.Config{})
	//
	//if err != nil {
	//	log.Fatal("Cannot connect to the Database")
	//	return
	//}
	//
	//gormDB = dbSession
}
