package student

import (
	"github.com/gin-gonic/gin"
	"library/pkg/auth"
	"library/pkg/librarian"
)

// LoginHandler godoc
// @Summary Login for students
// @Description get login to access resources
// @ID post-login-handler
// @Produce json
// @Success 200 {string, string} string, string "access_token", "refresh_token"
// @Router /login [post]

func SetupStudentRoutes(router *gin.Engine) {

	studentGroup := router.Group("")

	studentGroup.POST("/login", auth.LoginHandler)
	studentGroup.POST("/token_refresh", auth.RefreshHandler)

	studentGroup.POST("/register", registerStudentHandler)
	studentGroup.GET("/all_students", allStudentHandler)
	studentGroup.GET("/all_books", librarian.AllBooksHandler)
	studentGroup.POST("/borrow_book", bookBorrowHandler)
	studentGroup.GET("/get_borrowed_books", getBorrowedBookHandler)
	studentGroup.POST("/return_book", returnBookHandler)

}
