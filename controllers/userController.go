package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/hisyntax/crud-api/database"
	"github.com/hisyntax/crud-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, os.Getenv("USER_COLLECTION_NAME"))
var validate = validator.New()

//HashPassword is used to encrypt the password before it is stored in the mongo database
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

//VerifyPAssword chceks the input password while verifying it with the password in the mongo database
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "Incorrect Password Please try again"
		check = false
	}

	return check, msg
}

//CreateUSer is the api endpoint used to create a user
func SignUp(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate the user struct
	validateionErr := validate.Struct(user)
	if validateionErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validateionErr.Error()})
		return
	}

	//read through all the user email addresses
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		log.Panic(err)
		msg := "Error occured while checking for the Email"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	//check if the previously read user data already existes in the database
	//and if i does, throw an error but if it doesnt
	//then save it
	if count > 0 {
		msg := "This email address is aleady taken"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	//hash the user password provided
	password := HashPassword(*user.Password)
	//set the hased password to the user password in the User struct before saving in the database
	user.Password = &password

	//set a create and update time for the user
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	user.ID = primitive.NewObjectID()
	//assign the ID to the user_id
	user.User_id = user.ID.Hex()

	//insert the users data into the database
	retultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		msg := "User data was not created"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": retultInsertionNumber,
	})

}

//login is the api endpoint to signin the user inot the system
func Login(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	var foundUser models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		msg := "Incorrect email address"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
	if !passwordIsValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, foundUser)
}

//CreatePost is the api endpoint to create an item
func CreatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create post"})
}

//GetSinglePost is the api endpoint to create an item
func GetSinglePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get single post"})
}

//GetAllPost is the api endpoint to create an item
func GetAllPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get all post"})
}

//UpdatePost is the api endpoint to create an item
func UpdatePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Update post"})
}

//DeletePost is the api endpoint to create an item
func DeletePost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Delete post"})
}
