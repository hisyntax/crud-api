package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hisyntax/crud-api/controllers"
)

func UserRoutes() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/user/signup", controllers.SignUp)
		api.POST("/user/signin", controllers.Login)
	}

	router.Run(":" + port)
}
