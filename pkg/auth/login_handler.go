package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func LoginHandler(c *gin.Context) {

	var cred loginCred

	if err := c.ShouldBindWith(&cred, binding.JSON); err != nil {

		c.JSON(http.StatusUnauthorized, "Invalid credentials")
		return
	}

	validateLoginCred(&cred) // need to work on it

	isValid := isStudentInDB(cred)

	if isValid == false {

		c.JSON(http.StatusNotFound, fmt.Sprintf("Student with name %s not found", cred.Name))
		return

	}

	var t tokens
	t = generateTokens(cred)
	c.JSON(http.StatusCreated, t)
}
