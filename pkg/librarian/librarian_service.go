package librarian

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"library/database"
	"library/pkg/common"
	"net/http"
	"time"
)

func registerBook(newBook book, c *gin.Context) book {
	database.Connect()
	db := database.GetDBInstance()

	bookData := bson.D{
		{"ID", newBook.ID},
		{"Title", newBook.Title},
		{"Description", newBook.Description},
		{"Author", newBook.Author},
	}

	result, err := db.Collection("Books").InsertOne(context.TODO(), bookData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Book not registered in the library")
		return book{}
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Failed to get inserted document ID")
		return book{}
	}

	var insertedBook book

	err = db.Collection("Books").FindOne(context.TODO(), bson.M{"_id": insertedID}).Decode(&insertedBook)
	if err != nil {

		//fmt.Println("Type assertion failed. Actual type:", reflect.TypeOf(result.InsertedID))
		//fmt.Println("inserted id is:", result.InsertedID)
		c.JSON(http.StatusInternalServerError, "failed to get registered book from the library")
		return book{}
	}

	return insertedBook
}

func getAllBooks(c *gin.Context) []book {
	database.Connect()
	db := database.GetDBInstance()

	cursor, err := db.Collection("Books").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to get cursor")
		return nil
	}
	defer cursor.Close(context.TODO())

	var books []book
	if err := cursor.All(context.TODO(), &books); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to decode book data")
		return nil
	}

	return books
}

func getSingleBook(bookToCheck bookStatus, c *gin.Context) book {
	database.Connect()
	db := database.GetDBInstance()

	bookData := bson.M{
		"Title":  bookToCheck.Title,
		"Author": bookToCheck.Author,
	}

	var fetchedBook book
	result := db.Collection("Books").FindOne(context.TODO(), bookData)

	err := result.Decode(&fetchedBook)

	if err != nil {

		c.JSON(http.StatusInternalServerError, "failed to get book from the library, this book might not exist in the system")
		return book{}
	}
	return fetchedBook
}

func allBorrowRequests(c *gin.Context) []common.BorrowRequest {
	database.Connect()
	db := database.GetDBInstance()

	cursor, err := db.Collection("BookBorrowRequests").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to get cursor")
		return nil
	}
	defer cursor.Close(context.TODO())

	var bookBorrowRequests []common.BorrowRequest
	if err := cursor.All(context.TODO(), &bookBorrowRequests); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to decode borrow requests data")
		return nil
	}

	return bookBorrowRequests
}

func acceptBorrowRequest(c *gin.Context, acceptanceRequest borrowAcceptanceRequest) common.BorrowRequest {
	database.Connect()
	db := database.GetDBInstance()

	// task 1 --> update the IsBorrow status of book in the Book collection
	reqBook := bson.D{
		{"Title", acceptanceRequest.BookTitle},
	}
	update := bson.D{{"$set", bson.D{{"IsBorrowed", true}}}}

	_, err := db.Collection("Books").UpdateOne(context.TODO(), reqBook, update)
	//db.Client().StartSession()
	if err != nil {

		c.JSON(http.StatusInternalServerError, "failed to access book status")
		return common.BorrowRequest{}
	}

	// task 2 --> update the status of borrow request in the BookBorrowRequest collection
	req := bson.D{
		{"RequestID", acceptanceRequest.RequestID},
	}
	update = bson.D{{"$set", bson.D{{"Status", common.ISSUED}}}}
	var updatedRequest common.BorrowRequest

	err = db.Collection("BookBorrowRequests").FindOneAndUpdate(context.TODO(), req, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedRequest)

	if err != nil {
		//roll back the book status update to again false
		rollbackBookStatusUpdate := bson.D{{"$set", bson.D{{"IsBorrowed", false}}}}
		_, _ = db.Collection("Books").UpdateOne(context.TODO(), reqBook, rollbackBookStatusUpdate)

		c.JSON(http.StatusInternalServerError, "failed to access borrow request")
		return common.BorrowRequest{}
	}

	// task 3 --> insert the issued book in the BorrowedBooks collection
	//var borrowBook borrowedBook

	borrowedBookToDB := bson.D{
		{"RequestID", updatedRequest.RequestID},
		{"ID", updatedRequest.BookID},
		{"Title", updatedRequest.BookTitle},
		{"Author", updatedRequest.BookAuthor},
		{"BorrowerName", updatedRequest.BorrowerName},
		{"IssuedAt", time.Now().Local().Format(time.DateTime)},
	}

	_, err = db.Collection("BorrowedBooks").InsertOne(context.TODO(), borrowedBookToDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Borrowed Book has not recorded in the system")
		return common.BorrowRequest{}
	}

	return updatedRequest
}

func getBorrowedBooks(c *gin.Context) []BorrowedBook {
	database.Connect()
	db := database.GetDBInstance()

	cursor, err := db.Collection("BorrowedBooks").Find(context.TODO(), bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get cursor")

		return nil
	}

	var allBorrowedBooks []BorrowedBook
	if err = cursor.All(context.TODO(), &allBorrowedBooks); err != nil {
		c.JSON(http.StatusInternalServerError, "failed to decode borrow requests data")
		fmt.Println("Error:", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	return allBorrowedBooks

}
