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
	"github.com/hisyntax/crud-api/helpers"
	"github.com/hisyntax/crud-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var postCollection *mongo.Collection = database.OpenCollection(database.Client, os.Getenv("POST_COLLECTION_NAME"))
var userCollection *mongo.Collection = database.OpenCollection(database.Client, os.Getenv("USER_COLLECTION_NAME"))
var validate = validator.New()

//HashPassword is used to encrypt the password before it is stored in the mongo database
func HashPassword(password string) string {
	//encrypt the provided user password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

//VerifyPAssword chceks the input password while verifying it with the password in the mongo database
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	//check if the provided user password is the same as the already hashed password in the database
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
	//ope n a database conection to the mongo database
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	//close that connection after the resources in not in use
	defer cancel()

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate the user struct
	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
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
		msg := "this email already exists"
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
	user.User_id = user.ID.Hex()

	//generate a token for the user on signup
	token, refreshToken, _ := helpers.GenerateAllTokens(*user.Email, user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

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
	//ope n a database conection to the mongo database
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	//close that connection after the resources in not in use
	defer cancel()

	var user models.User
	var foundUser models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if the user email is already taken
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		msg := "Incorrect email"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	//chack if the provided password is the same as the password in the database
	passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)
	if !passwordIsValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := helpers.GenerateAllTokens(*foundUser.Email, foundUser.User_id)

	helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)

	c.JSON(http.StatusOK, foundUser)
	// c.Redirect(500, "/api/v1/post/create")
}
