package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MiddlewaresSaldo(c *gin.Context){

	saldo := c.PostForm("saldo")
	if saldo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"campo saldo obligatorio",
		})
		c.Abort()
		return
	}
	saldoint,err := strconv.ParseFloat(saldo, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"saldo tiene que ser numerico",
		})
		c.Abort()
		return
	}
	if saldoint < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"saldo no puede ser un numero negativo",
		})
		c.Abort()
		return
	}

	c.Next()
}