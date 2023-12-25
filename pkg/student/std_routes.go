package student

import (
	"github.com/gin-gonic/gin"
)

func SetupStudentRoutes(router *gin.Engine) {

	studentGroup := router.Group("/student")

	studentGroup.POST("/register", registerStudentHandler)
	studentGroup.GET("/allStudents", allStudentHandler)

}
