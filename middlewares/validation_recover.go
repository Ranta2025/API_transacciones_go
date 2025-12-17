package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validation_recover(c *gin.Context){
	gmail := c.Query("gmail")
	if gmail == "" {
		c.JSON(http.StatusBadRequest, "campo gmail obligatorio")
		c.Abort()
		return
	}
	c.Next()
}

func Validation_recover_post(c *gin.Context){
	username := c.PostForm("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, "campo username obligatorio")
		c.Abort()
		return
	}

	codigo := c.PostForm("codigo")
	if codigo == "" {
		c.JSON(http.StatusBadRequest, "campo codigo obligatorio")
		c.Abort()
		return
	}
	c.Next()
}