package librarian

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"library/database"
	"library/pkg/common"
	"net/http"
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
		//fmt.Println("Type assertion failed. Actual type:", reflect.TypeOf(result.InsertedID))
		//fmt.Println("inserted id is:", result.InsertedID)
		c.JSON(http.StatusInternalServerError, "failed to get book from the library")
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
