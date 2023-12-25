package main

//@Summary Get list of students
//@Description Get a list of all students
//@ID get-all-students
//@Produce json
//@Success 200 {array} Student
//@Router /students [get]

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"library/docs"
	"library/pkg/auth"
	"library/pkg/student"
	"log"
	"net/http"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//cmd.Execute()
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}

	//swagger := router.Group("Swagger")
	{
		docs.SwaggerInfo.Title = "Library"
		docs.SwaggerInfo.Description = "Library and Students"
		docs.SwaggerInfo.Version = "1"
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//	swagger.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}

	//router.NoRoute(func(c *gin.Context) {
	//	c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	//})

	//router.POST()

	student.SetupStudentRoutes(router)

	router.POST("/login", auth.LoginHandler)

	router.Run("localhost:8000")
}
