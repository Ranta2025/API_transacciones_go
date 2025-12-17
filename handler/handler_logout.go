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

func Handler_logout(data *sql.DB, rd *redis.Client) gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		user := models.UserLogout{
			Username: c.PostForm("username"),
			Gmail: c.PostForm("gmail"),
			Password: utils.Hash_password(c.PostForm("password")),
		}
		str := db.Logout(ctx, user,rd , data)
		if str != "" {
			c.String(http.StatusBadRequest, str)
			return 
		}
		c.JSON(200, "usuario creado")
	}
}