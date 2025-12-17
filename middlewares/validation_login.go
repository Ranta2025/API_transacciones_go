package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validation_Login(c *gin.Context){
	password := c.PostForm("password")
	username := c.PostForm("username")

	if username == "" {
		c.String(http.StatusBadRequest, "campo usuario obligatorio")
		c.Abort()
		return 
	}

	if password == "" {
		c.String(http.StatusBadRequest, "campo password obligatorio")
		c.Abort()
		return 
	}

	c.Next()
}