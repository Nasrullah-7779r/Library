package librarian

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"library/pkg/auth"
	"library/pkg/common"
	"net/http"
)

// loginHandler godoc
// @Summary login for Librarian
// @Description get Librarian logged in
// @ID login
// @Produce json
// @Accept x-www-form-urlencoded
// @Param username formData string true "Librarian's username"
// @Param password formData string true "Librarian's password" format(password)
// @Success 201 {object}  auth.Tokens
// @Tags Auth
// @Router /librarian/login [post]
func loginHandler(c *gin.Context) {
	var cred common.LoginCred

	if err := c.ShouldBindWith(&cred, binding.Form); err != nil {
		c.JSON(http.StatusUnauthorized, "Invalid credentials")
		return
	}
	if cred.Name == "librarian" && cred.Password == "123" {
		var t auth.Tokens
		t = auth.GenerateTokens(cred)
		c.JSON(http.StatusCreated, t)
		return

	}

	c.JSON(http.StatusUnauthorized, "invalid credentials")
}

// refreshHandler godoc
// @Summary refreshToken
// @Description get new access token to access resources, if expired
// @ID Librarian refresh-handler
// @Produce json
// @Param refresh_token header string true "Librarian's refresh token"
// @Success 201 {object}  auth.AccessToken
// @Tags Auth
// @Router /librarian/token_refresh [post]
func refreshHandler(c *gin.Context) {
	//refreshTokenHeader := c.GetHeader("Authorization")
	//
	//refreshToken := strings.TrimPrefix(refreshTokenHeader, "Bearer ")

	refreshToken, err := auth.GetTokenFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	var user common.LoginCred
	var isVerified bool

	isVerified, user = auth.VerifyRefreshToken(refreshToken)
	if !isVerified {
		c.JSON(http.StatusUnprocessableEntity, "Invalid token")
		return
	}

	var t auth.AccessToken
	t = auth.GenerateAccessToken(user)

	c.JSON(http.StatusCreated, t)
}

// registerBookHandler godoc
// @Summary registration of new Book
// @Description Librarian can register new Book
// @ID register
// @Produce json
// @Param id body string true "Book ID"
// @Param title body string true "Book Title"
// @Param description body string true "Book Description"
// @Param author body string true "Book Author"
// @Success 201 {object} book
// @Tags Librarian
// @Router /librarian/register_book [post]
func registerBookHandler(c *gin.Context) {

	var newBook book

	//var accessToken string
	//
	//tokenHeader := c.GetHeader("Authorization")
	//
	//accessToken = strings.TrimPrefix(tokenHeader, "Bearer ")

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

	if err := c.ShouldBindWith(&newBook, binding.JSON); err != nil {
		//fmt.Println("book data in bind:", newBook)
		c.JSON(http.StatusBadRequest, "invalid book data")
		return
	}
	//fmt.Println("book data:", newBook)

	insertedBook := registerBook(newBook, c)

	c.JSON(http.StatusCreated, insertedBook)
}

// AllBooksHandler godoc
// @Summary all books
// @Description get all books of library
// @ID get all books
// @Produce json
// @Success 200 {object} []book
// @Tags Librarian
// @Router /librarian/all_books [get]
func AllBooksHandler(c *gin.Context) {

	//var accessToken string
	//
	//tokenHeader := c.GetHeader("Authorization")
	//
	//accessToken = strings.TrimPrefix(tokenHeader, "Bearer ")

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

	allBooks := getAllBooks(c)

	c.JSON(http.StatusOK, allBooks)

}

// SingleBookHandler godoc
// @Summary get single book
// @Description
// @ID get single book
// @Produce json
// @Success 200 {object} book
// @Tags Librarian
// @Router /librarian/book [get]
func singleBookHandler(c *gin.Context) {

	//var accessToken string
	//
	//tokenHeader := c.GetHeader("Authorization")
	//
	//accessToken = strings.TrimPrefix(tokenHeader, "Bearer ")

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

	var bookToCheck bookStatus

	if err := c.ShouldBindWith(&bookToCheck, binding.JSON); err != nil {
		//fmt.Println("book data in bind:", newBook)
		c.JSON(http.StatusBadRequest, "invalid book data")
		return
	}

	respectiveBook := getSingleBook(bookToCheck, c)

	c.JSON(http.StatusOK, respectiveBook)

}

// allBorrowRequestsHandler godoc
// @Summary all books borrow requests
// @Description
// @ID
// @Produce json
// @Success 200 {object} []common.BorrowRequest
// @Tags Librarian
// @Router /librarian/borrow_requests [get]
func allBorrowRequestsHandler(c *gin.Context) {

	//var accessToken string
	//
	//tokenHeader := c.GetHeader("Authorization")
	//
	//accessToken = strings.TrimPrefix(tokenHeader, "Bearer ")

	accessToken, err := auth.GetTokenFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}

	isVerified, _ := auth.VerifyAccessToken(accessToken)
	if !isVerified {
		c.JSON(http.StatusUnprocessableEntity, "Invalid token")
		return
	}

	allBorrowRequest := allBorrowRequests(c)

	c.JSON(http.StatusOK, allBorrowRequest)

}

func acceptBorrowRequestHandler(c *gin.Context) {

	accessToken, err := auth.GetTokenFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}

	isVerified, _ := auth.VerifyAccessToken(accessToken)
	if !isVerified {
		c.JSON(http.StatusUnprocessableEntity, "Invalid token")
		return
	}

	var acceptanceRequest borrowAcceptanceRequest

	c.ShouldBindWith(&acceptanceRequest, binding.JSON)

	acceptedBorrowRequest := acceptBorrowRequest(c, acceptanceRequest)

	c.JSON(http.StatusAccepted, acceptedBorrowRequest)

}

func getBorrowedBooksHandler(c *gin.Context) {

	accessToken, err := auth.GetTokenFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
	}

	isVerified, _ := auth.VerifyAccessToken(accessToken)
	if !isVerified {
		c.JSON(http.StatusUnprocessableEntity, "Invalid token")
		return
	}

	//c.ShouldBindWith(&acceptanceRequest, binding.JSON)

	allBorrowedBooks := getBorrowedBooks(c)

	c.JSON(http.StatusOK, allBorrowedBooks)
}
