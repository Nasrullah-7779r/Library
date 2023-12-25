package student

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/pkg/auth"
	"net/http"
)

func HelloStudent(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello student",
	})

}

func registerStudentHandler(c *gin.Context) {

	var student Student

	if err := c.BindJSON(&student); err != nil {

		c.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	insertedStudent := addStudentInDB(&student, c)

	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, "Failed to add student to the database")
	//	return
	//}

	c.JSON(http.StatusCreated, insertedStudent)

}

func allStudentHandler(c *gin.Context) {
	var body string
	var accessToken string

	if err := c.BindJSON(&body); err != nil {

		fmt.Println("body is:", string(body))
		c.JSON(http.StatusBadRequest, "Invalid data")
		return
	}

	fmt.Println("after bind body is:", string(body))

	isVerified := auth.VerifyAccessToken(accessToken)
	if isVerified == false {
		c.JSON(http.StatusUnprocessableEntity, "Invalid token")
		return
	}

	//students := getAllStudents(c)

	c.JSON(http.StatusOK, "Students")
}
