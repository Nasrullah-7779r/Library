package librarian

import (
	"github.com/gin-gonic/gin"
)

func SetupLibrarianRoutes(router *gin.Engine) {

	librarianGroup := router.Group("librarian")

	librarianGroup.POST("/login", loginHandler)
	librarianGroup.POST("/token_refresh", refreshHandler)

	librarianGroup.POST("/register_book", registerBookHandler)
	librarianGroup.GET("/all_books", AllBooksHandler)
	librarianGroup.GET("/borrow_requests", allBorrowRequestsHandler)
	librarianGroup.GET("/book", singleBookHandler)
	librarianGroup.POST("/accept_borrow_request", acceptBorrowRequestHandler)
	librarianGroup.GET("/get_borrowed_books", getBorrowedBooksHandler)

}
