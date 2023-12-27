package student

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"library/database"
	"library/pkg/common"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func addStudentInDB(newStudent *common.Student, c *gin.Context) common.StudentOut {

	database.Connect()
	db := database.GetDBInstance()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newStudent.Password), 12)
	if err != nil {
		fmt.Errorf("hashpassword error")
		return common.StudentOut{}
	}
	std := bson.D{
		{"Name", newStudent.Name},
		{"Email", newStudent.Email},
		{"Password", hashedPassword},
	}

	result, err := db.Collection("Students").InsertOne(context.TODO(), std)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Student not added in the db")
		return common.StudentOut{}
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Failed to get inserted document ID")
		return common.StudentOut{}
	}

	var insertedStudent common.StudentOut

	err = db.Collection("Students").FindOne(context.TODO(), bson.M{"_id": insertedID}).Decode(&insertedStudent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to fetch inserted student from the db")
		return common.StudentOut{}
	}

	return insertedStudent

}

func getAllStudents(c *gin.Context) []common.StudentOut {
	database.Connect()
	db := database.GetDBInstance()

	cursor, err := db.Collection("Students").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to get cursor")
		return nil
	}
	defer cursor.Close(context.TODO())

	var students []common.StudentOut
	if err := cursor.All(context.TODO(), &students); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to decode student data")
		return nil
	}

	return students
}

func borrowAbook(borrowRequest common.BorrowRequest, c *gin.Context) common.BorrowRequest {
	database.Connect()
	db := database.GetDBInstance()

	request := bson.D{
		{"RequestID", strconv.Itoa(rand.Intn(90000) + 10000)},
		{"BookTitle", borrowRequest.BookTitle},
		{"BookAuthor", borrowRequest.BookAuthor},
		{"BorrowerName", borrowRequest.BorrowerName},
		{"Status", common.PENDING},
		{"Time", time.Now()},
	}

	result, err := db.Collection("BookBorrowRequests").InsertOne(context.TODO(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "borrow request not submitted")
		return common.BorrowRequest{}
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, "failed to get inserted document ID")
		return common.BorrowRequest{}
	}
	fmt.Println("request:", request)
	var insertedBorrowRequest common.BorrowRequest

	err = db.Collection("BookBorrowRequests").FindOne(context.TODO(), bson.M{"_id": insertedID}).Decode(&insertedBorrowRequest)
	if err != nil {
		fmt.Println("inserted id:", insertedID)
		c.JSON(http.StatusInternalServerError, "failed to fetch submitted request from the db")
		return common.BorrowRequest{}
	}
	return insertedBorrowRequest
}
