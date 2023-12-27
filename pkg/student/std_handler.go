package student

import (
	"github.com/gin-gonic/gin"
	"library/pkg/common"
	"library/pkg/student/std_auth"
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
	if err := c.BindJSON(&cred); err != nil {

		c.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	//var isVerified bool
	//isVerified, cred = auth.VerifyAccessToken(accessToken)
	//if !isVerified {
	//	c.JSON(http.StatusUnprocessableEntity, "Invalid token")
	//	return
	//}

	isVerified := std_auth.IsStudentInDB(cred)

	if !isVerified {
		c.JSON(http.StatusNotFound, "Student not found")
		return
	}
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
func BookBorrowHandler(c *gin.Context) {
	var request common.BorrowRequest
	if err := c.BindJSON(&request); err != nil {

		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	borrowRequest := borrowAbook(request, c)

	c.JSON(http.StatusCreated, borrowRequest)
}
