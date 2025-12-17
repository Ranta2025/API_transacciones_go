package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidationUser(c *gin.Context){
	rol, exist := c.Get("rol")
	if !exist {
		c.String(400, "error")
		return 
	}
	rolstr := rol.(string)

	if rolstr != "user" {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"error":"acceso restringido",
		})
		c.Abort()
		return
	}
	c.Next()
}

func ValidationAdmin(c *gin.Context){
	rol, exist := c.Get("rol")
	if !exist {
		c.String(400, "error")
		return 
	}
	rolstr := rol.(string)

	if rolstr != "admin" {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"error":"acceso restringido",
		})
		c.Abort()
		return
	}
	c.Next()
}