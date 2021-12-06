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
	//ope n a database conection to the mongo database
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	//close that connection after the resources in not in use
	defer cancel()

	var post models.Post
	// var user models.User
	// var findUser models.User

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
		"title":   post.Title,
		"user_id": retultInsertionNumber,
	})

}

//GetSinglePost is the api endpoint to create an item
func GetSinglePost(c *gin.Context) {
	// //ope n a database conection to the mongo database
	// var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	// //close that connection after the resources in not in use
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
	//ope n a database conection to the mongo database
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	//close that connection after the resources in not in use
	defer cancel()
	var post []models.Post

	//query through the database to fine all available posts
	cusor, err := postCollection.Find(ctx, bson.D{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong Please try again sometime!"})
		return
	}
	//pass the results gotten as its own individual results
	err = cusor.All(ctx, &post)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer cusor.Close(ctx)
	if err := cusor.Err(); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"post": post})
}

//UpdatePost is the api endpoint to create an item
func UpdatePost(c *gin.Context) {
	// //ope n a database conection to the mongo database
	// var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	// //close that connection after the resources in not in use
	// defer cancel()
	c.JSON(http.StatusOK, gin.H{"message": "Update post"})
}

//DeletePost is the api endpoint to create an item
func DeletePost(c *gin.Context) {
	//ope n a database conection to the mongo database
	// var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	// //close that connection after the resources in not in use
	// defer cancel()
	c.JSON(http.StatusOK, gin.H{"message": "Delete post"})
}
