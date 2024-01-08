package student

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"library/database"
	"library/pkg/common"
	"library/pkg/librarian"
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

	filter := bson.D{{"Title", borrowRequest.BookTitle}}

	type bookID struct {
		ID uint `json:"id"`
	}
	var resultID bookID

	err := db.Collection("Books").FindOne(context.TODO(), filter).Decode(&resultID)
	if err != nil {
		fmt.Println("Error", err)
		c.JSON(http.StatusPartialContent, fmt.Sprintf("error %v", err))
		return common.BorrowRequest{}
	}

	request := bson.D{
		{"RequestID", strconv.Itoa(rand.Intn(90000) + 10000)},
		{"BookID", resultID.ID},
		{"BookTitle", borrowRequest.BookTitle},
		{"BookAuthor", borrowRequest.BookAuthor},
		{"BorrowerName", borrowRequest.BorrowerName},
		{"Status", common.PENDING},
		{"CreatedAt", time.Now().Local().Format(time.DateTime)},
	}

	result, errr := db.Collection("BookBorrowRequests").InsertOne(context.TODO(), request)
	if errr != nil {
		c.JSON(http.StatusInternalServerError, "borrow request not submitted")
		return common.BorrowRequest{}
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, "failed to get inserted document ID")
		return common.BorrowRequest{}
	}

	var insertedBorrowRequest common.BorrowRequest

	err = db.Collection("BookBorrowRequests").FindOne(context.TODO(), bson.M{"_id": insertedID}).Decode(&insertedBorrowRequest)
	if err != nil {

		c.JSON(http.StatusInternalServerError, "failed to fetch submitted request from the db")
		return common.BorrowRequest{}
	}
	return insertedBorrowRequest
}

func getBorrowedBook(stdName string, c *gin.Context) []librarian.BorrowedBook {

	database.Connect()
	db := database.GetDBInstance()

	// to query in case insensitive mode
	filter := bson.D{{"BorrowerName", primitive.Regex{Pattern: stdName, Options: "i"}}}

	cursor, err := db.Collection("BorrowedBooks").Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error", err)
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("error %v", err))
		return nil
	}

	var borrowedBooks []librarian.BorrowedBook
	if err := cursor.All(context.TODO(), &borrowedBooks); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to decode borrowed book data")
		return nil
	}
	defer cursor.Close(context.TODO())

	return borrowedBooks
}

func returnBook(requestID string, bookID uint, c *gin.Context) string {

	database.Connect()
	db := database.GetDBInstance()
	fmt.Print("request id: ", requestID, "book id:", bookID)
	// task 1 --> update the request status of Borrow Request in the Book collection
	req := bson.D{
		{"RequestID", requestID},
	}

	update := bson.D{{"$set", bson.D{{"Status", common.COMPLETED}}}}
	var uResult *mongo.UpdateResult
	uResult, err := db.Collection("BookBorrowRequests").UpdateOne(context.TODO(), req, update)

	if err != nil {

		c.AbortWithStatus(http.StatusInternalServerError)
		return "failed to update borrow request status"
	}

	if uResult.MatchedCount == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return "borrow request not found"
	}

	if uResult.MatchedCount > 0 && uResult.ModifiedCount == 0 {
		c.AbortWithStatus(http.StatusOK) // or http.StatusNoContent depending on your needs
		return "borrow request is already completed"
	}

	// task 2 --> update the book isBorrow status in DB
	bookToUpdate := bson.D{
		{"ID", bookID},
	}

	update = bson.D{{"$set", bson.D{{"IsBorrowed", false}}}}

	_, err = db.Collection("Books").UpdateOne(context.TODO(), bookToUpdate, update)

	if err != nil {

		c.AbortWithStatus(http.StatusInternalServerError)
		return "failed to update book isBorrow status"
	}

	// task 3 --> remove that specific borrowed book which is being returned
	bookRecordToRemove := bson.D{
		{"RequestID", requestID},
	}

	_, err = db.Collection("BorrowedBooks").DeleteOne(context.TODO(), bookRecordToRemove)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return "Book has been returned successfully"
}
