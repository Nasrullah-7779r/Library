package std_auth

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"library/database"
	"library/pkg/common"
	"log"
)

func IsStudentInDB(userCred common.LoginCred) bool {

	database.Connect()
	db := database.GetDBInstance()

	std := bson.D{
		// to perform case-insensitive search for name
		{"Name", bson.D{{"$regex", primitive.Regex{Pattern: userCred.Name, Options: "i"}}}},
	}

	result := db.Collection("Students").FindOne(context.TODO(), std)
	if result.Err() == mongo.ErrNoDocuments {
		fmt.Println("after no document")
		return false
	}

	var fetchedStudent common.Student
	if err := result.Decode(&fetchedStudent); err != nil {
		log.Print("decoding error")
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(fetchedStudent.Password), []byte(userCred.Password))
	if err != nil {
		log.Print("password did not matched")
		return false
	}

	return true
}
