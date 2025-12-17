package handler

import (
	"database/sql"
	"proyecto_transacciones/db"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Handler_Recover(data *sql.DB, rd *redis.Client) gin.HandlerFunc{
	return func (c *gin.Context)  {
		ctx := c.Request.Context()
		gmail := c.Query("gmail")
		
		flag, mensaje := db.Check_gmail_invalid(ctx, gmail, rd, data)
		if !flag {
			c.JSON(200, gin.H{
				"mensaje": mensaje,
			})
			return
		}
		c.JSON(200, gin.H{
			"codigo": mensaje,
		})
	}
}