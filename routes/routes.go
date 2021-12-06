package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hisyntax/crud-api/controllers"
	"github.com/hisyntax/crud-api/middleware"
)

func UserRoutes() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := gin.Default()

	//user routes
	user := router.Group("/api/v1")
	{
		user.POST("/user/signup", controllers.SignUp)
		user.POST("/user/signin", controllers.Login)
	}

	//post routes
	post := router.Group("/api/v1")
	post.Use(middleware.Authentication)
	{
		post.POST("/post/create", controllers.CreatePost)
		post.GET("/post/:post_id", controllers.GetSinglePost)
		post.GET("/post/posts", controllers.GetAllPost)
		post.PATCH("post/:post_id", controllers.UpdatePost)
		post.DELETE("/post/:post_id", controllers.DeletePost)
		post.GET("/user/test", controllers.Test)
	}

	router.Run(":" + port)
}
