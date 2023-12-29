package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"library/pkg/common"
	"library/pkg/student/std_auth"
	"net/http"
)

// LoginHandler godoc
// @Summary Login for students
// @Description get login to access resources
// @ID post-login-handler
// @Produce json
// @Success 201 {obj}  "access_token", "refresh_token"
// @Tags Auth
// @Router /login [post]
func LoginHandler(c *gin.Context) {

	var cred common.LoginCred

	if err := c.ShouldBindWith(&cred, binding.Form); err != nil {
		fmt.Println("cred are", cred)
		c.JSON(http.StatusUnauthorized, "Invalid credentials")
		return
	}

	validateLoginCred(&cred) // need to work on it

	isValid := std_auth.IsStudentInDB(cred)

	if isValid == false {

		c.JSON(http.StatusNotFound, fmt.Sprintf("Student with name %s not found", cred.Name))
		return

	}

	var t Tokens
	t = GenerateTokens(cred)
	c.JSON(http.StatusCreated, t)
}

// RefreshHandler godoc
// @Summary refreshToken
// @Description get new access token to access resources, if expired
// @ID student refresh-handler
// @Produce json
// @Success 201 {object}  AccessToken
// @Tags Auth
// @Router /token_refresh [post]
func RefreshHandler(c *gin.Context) {

	//var cred loginCred
	//
	//if err := c.ShouldBindWith(&cred, binding.JSON); err != nil {
	//
	//	c.JSON(http.StatusUnauthorized, "Invalid credentials")
	//	return
	//}
	//
	//isValid := isStudentInDB(cred)
	//
	//if isValid == false {
	//
	//	c.JSON(http.StatusNotFound, fmt.Sprintf("Student with name %s not found", cred.Name))
	//	return
	//
	//}

	//refreshTokenHeader := c.GetHeader("Authorization")
	//
	//refreshToken := strings.TrimPrefix(refreshTokenHeader, "Bearer ")

	refreshToken, err := GetTokenFromRequest(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	var user common.LoginCred
	var isVerified bool
	isVerified, user = VerifyRefreshToken(refreshToken)
	if !isVerified {
		c.JSON(http.StatusUnprocessableEntity, "Invalid token")
		return
	}

	var t AccessToken
	t = GenerateAccessToken(user)

	c.JSON(http.StatusCreated, t)
}
