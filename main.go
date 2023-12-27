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
	"library/pkg/librarian"
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

	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"}, // You might want to restrict this to specific origins in production
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	students := v1.Group("/students")
	{
		students.GET("/")

	}

	{
		docs.SwaggerInfo.Title = "Library"
		docs.SwaggerInfo.Description = "Library and Students"
		docs.SwaggerInfo.Version = "1"
		docs.SwaggerInfo.Host = "localhost:8000"
		docs.SwaggerInfo.BasePath = "/"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	}

	student.SetupStudentRoutes(router)
	librarian.SetupLibrarianRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8000")
}

//package main
//
////// @BasePath /api/v1
////
////// PingExample godoc
////// @Summary ping example
////// @Schemes
////// @Description do ping
////// @Tags example
////// @Accept json
////// @Produce json
////// @Success 200 {string} Helloworld
////// @Router /example/helloworld [get]
////func Helloworld(g *gin.Context) {
////	g.JSON(http.StatusOK, "helloworld")
////}
////
////func main() {
////	r := gin.Default()
////	docs.SwaggerInfo.BasePath = "/api/v1"
////	v1 := r.Group("/api/v1")
////	{
////		eg := v1.Group("/example")
////		{
////			eg.GET("/helloworld", Helloworld)
////		}
////	}
////
////	{
////		docs.SwaggerInfo.Title = "Library"
////		docs.SwaggerInfo.Description = "Library and Students"
////		docs.SwaggerInfo.Version = "1"
////	}
////	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
////	r.Run(":8080")
////
////}
