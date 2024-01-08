package student

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"library/pkg/auth"
	"library/pkg/common"
	"net/http"
)

func HelloStudent(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello student",
	})

}

// registerStudentHandler godoc
// @Summary register new student
// @Description student can get him/herself registered with data
// @ID post-register-student-handler
// @Produce json
// @Success 201 {object}  common.Student
// @Tags Student
// @Router /register [post]
func registerStudentHandler(c *gin.Context) {

	var student common.Student

	if err := c.BindJSON(&student); err != nil {

		c.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	insertedStudent := addStudentInDB(&student, c)

	c.JSON(http.StatusCreated, insertedStudent)
}

// allStudentHandler godoc
// @Summary get all registered students
// @Description get all registered students
// @ID get-get-all-student-handler
// @Produce json
// @Success 200 {object} []common.Student
// @Tags Student
// @Router /all_students [get]
func allStudentHandler(c *gin.Context) {

	//var accessToken string

	//tokenHeader := c.GetHeader("Authorization")
	//
	//accessToken = strings.TrimPrefix(tokenHeader, "Bearer ")

	var cred common.LoginCred
	if err := c.ShouldBindWith(&cred, binding.Form); err != nil {

		c.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	//var isVerified bool
	//isVerified, cred = auth.VerifyAccessToken(accessToken)
	//if !isVerified {
	//	c.JSON(http.StatusUnprocessableEntity, "Invalid token")
	//	return
	//}

	//isVerified := std_auth.IsStudentInDB(cred)
	//
	//if !isVerified {
	//	c.JSON(http.StatusNotFound, "Student not found")
	//	return
	//}
	students := getAllStudents(c)

	c.JSON(http.StatusOK, students)
}

// BookBorrowHandler godoc
// @Summary put a borrow request to acquire a book
// @Description BookBorrowHandler
// @ID borrow-a-book
// @Produce json
// @Success 200 {object} common.BorrowRequest
// @Tags Student
// @Router /borrow_book [get]
func bookBorrowHandler(c *gin.Context) {

	accessToken, err := auth.GetTokenFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	isVerified, name := auth.VerifyAccessToken(accessToken)
	if !isVerified {
		c.JSON(http.StatusUnprocessableEntity, "Invalid token")
		return
	}

	var request common.BorrowRequest
	if err := c.BindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}
	caser := cases.Title(language.English)
	request.BorrowerName = caser.String(name)

	borrowRequest := borrowAbook(request, c)

	c.JSON(http.StatusCreated, borrowRequest)
}

func getBorrowedBookHandler(c *gin.Context) {
	accessToken, err := auth.GetTokenFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	isVerified, name := auth.VerifyAccessToken(accessToken)
	if !isVerified {
		c.JSON(http.StatusUnprocessableEntity, "Invalid token")
		return
	}

	borrowedBooks := getBorrowedBook(name, c)

	c.JSON(http.StatusOK, borrowedBooks)

}

func returnBookHandler(c *gin.Context) {
	accessToken, err := auth.GetTokenFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	isVerified, _ := auth.VerifyAccessToken(accessToken)
	if !isVerified {
		c.JSON(http.StatusUnprocessableEntity, "Invalid token")
		return
	}
	type returnBookData struct {
		RequestID string `json:"request_id" binding:"required"`
		BookID    uint   `json:"book_id" binding:"required"`
	}
	var returnBookModel returnBookData

	if err = c.BindJSON(&returnBookModel); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println("request id: ", returnBookModel.RequestID, "book id:", returnBookModel.BookID)

	result := returnBook(returnBookModel.RequestID, returnBookModel.BookID, c)

	c.JSON(http.StatusAccepted, map[string]string{"message": result})

}
