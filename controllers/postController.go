package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hisyntax/crud-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreatePost is the api endpoint to create an item
func CreatePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var post models.Post
	var user models.User
	// var findUser models.User

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := c.Request.Cookie("auth")
	if err == http.ErrNoCookie {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not signed in"})
		return
	} else {
		//set a create and update time for the user
		post.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		post.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		post.ID = primitive.NewObjectID()
		//assign the ID to the user_id
		post.Post_id = post.ID.Hex()

		//insert the users data into the database
		retultInsertionNumber, insertErr := postCollection.InsertOne(ctx, post)
		if insertErr != nil {
			msg := "User data was not created"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":    user.Email,
			"title":   post.Title,
			"user_id": retultInsertionNumber,
		})
	}

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}

//GetSinglePost is the api endpoint to create an item
func GetSinglePost(c *gin.Context) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// var post []models.Post
	// params := c.Param("post_id")
	// query, err := postCollection.Find(ctx, post)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// c.JSON(http.StatusOK, gin.H{"message": query})

}

//GetAllPost is the api endpoint to create an item
func GetAllPost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var post models.Post
	query, getUsersErr := postCollection.Find(ctx, bson.M{"title": post.Title})
	if getUsersErr != nil {
		log.Fatal(getUsersErr)
	}

	c.JSON(http.StatusOK, gin.H{"message": query})
}

//UpdatePost is the api endpoint to create an item
func UpdatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update post"})
}

//DeletePost is the api endpoint to create an item
func DeletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete post"})
}
