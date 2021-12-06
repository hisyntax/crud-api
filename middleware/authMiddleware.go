package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hisyntax/crud-api/helpers"
)

//authentication validates the token and authorizes the user

func Authentication(c *gin.Context) {
	//specify the header key to hold the value which is the token for the user
	clientToken := c.Request.Header.Get("token")

	//check if the token is provided
	if clientToken == "" {
		msg := "Client token not provided"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(clientToken)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}

	c.Set("email", claims.Email)
	c.Set("uid", claims.Uid)

	c.Next()

}
