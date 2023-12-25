package student

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"library/database"
	"net/http"
)

func addStudentInDB(newStudent *Student, c *gin.Context) Student {

	database.Connect()
	db := database.GetDBInstance()
	std := bson.D{
		{"Name", newStudent.Name},
		{"Email", newStudent.Email},
		{"Password", newStudent.Password},
	}

	result, err := db.Collection("Students").InsertOne(context.TODO(), std)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Student not added in the db")
		return Student{}
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Failed to get inserted document ID")
		return Student{}
	}

	var insertedStudent Student

	err = db.Collection("Students").FindOne(context.TODO(), bson.M{"_id": insertedID}).Decode(&insertedStudent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to fetch inserted student from the db")
		return Student{}
	}

	return insertedStudent

}

func getAllStudents(c *gin.Context) []Student {
	database.Connect()
	db := database.GetDBInstance()

	cursor, err := db.Collection("Students").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to get cursor")
		return nil
	}
	defer cursor.Close(context.TODO())

	var students []Student
	if err := cursor.All(context.TODO(), &students); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to decode student data")
		return nil
	}

	return students
}
