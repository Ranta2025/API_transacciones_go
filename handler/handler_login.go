package handler

import (
	"database/sql"
	"net/http"
	"proyecto_transacciones/db"
	"proyecto_transacciones/models"
	"proyecto_transacciones/utils"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Handler_login(data *sql.DB, rd *redis.Client) gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		user := models.Login{
			Username: c.PostForm("username"),
			Password: c.PostForm("password"),
		}
		flag := db.Logindb(ctx, data, rd, user)
		if  flag != ""{
			c.JSON(http.StatusBadRequest, gin.H{
				"error":flag,
			})
			return 
		}
		rol := db.GetRol(c, user.Username, data)
		token, err :=utils.Generated_token(user.Username, rol)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"Problema al generar el token":"vuelva a inicial sesion",
			})
			return
		}

		c.JSON(200, gin.H{
			"credenciales validas": token,
		})
	}
}