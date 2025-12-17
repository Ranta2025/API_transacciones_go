package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewaresTransfer(c *gin.Context) {

	saldo := c.PostForm("username")
	if saldo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "campo username obligatorio",
		})
		c.Abort()
		return
	}

	c.Next()
}