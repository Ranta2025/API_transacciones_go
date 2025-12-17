package handler

import (
	"database/sql"
	"net/http"
	"proyecto_transacciones/db"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Recover_post(data *sql.DB, rd *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := ctx.Request.Context()
		username := ctx.PostForm("username")
		codigo := ctx.PostForm("codigo")

		flag, mensaje := db.Recoverdb(c ,username ,codigo ,rd ,data)

		if !flag {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": mensaje,
			})
			return 
		}

		ctx.JSON(200, gin.H{
			"valido" : mensaje,
		})
	}
}