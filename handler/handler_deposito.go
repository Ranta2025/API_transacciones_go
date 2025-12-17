package handler

import (
	"database/sql"
	"net/http"
	"proyecto_transacciones/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func HandlerDeposito(data *sql.DB, rd *redis.Client) gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		username, exist := c.Get("username")
		if !exist {
			c.JSON(400, gin.H{
				"error" : "no se tiene usuario",
			})
			return 
		}

		userstr := username.(string)

		saldo, err := strconv.ParseFloat(c.PostForm("saldo"), 64)

		mensaje, err := db.DepositarDB(ctx, userstr, saldo, data, rd)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error" : err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"exito": mensaje,
		})
	}
}