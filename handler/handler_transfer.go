package handler

import (
	"database/sql"
	"net/http"
	"proyecto_transacciones/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func Handler_user_getSaldo(data *sql.DB, rd *redis.Client) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		c := ctx.Request.Context()
		username, exist := ctx.Get("username")
		if !exist {
			ctx.String(400, "ja")
		}
		userstr := username.(string)
		saldo := db.Check_saldo(&c, userstr, data, rd)
		ctx.JSON(200, gin.H{
			"saldo":saldo,
		})
	}
}
func HandlerUserTransferSaldo(data *sql.DB, rd *redis.Client) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		c := ctx.Request.Context()
		username, exist := ctx.Get("username")
		if !exist {
			ctx.String(400, "error")
			return 
		}
		userstr := username.(string)
		saldo := ctx.PostForm("saldo")
		saldofloat,_ := strconv.ParseFloat(saldo, 64)
		user := ctx.PostForm("username")
		err := db.TranderSaldo(c, userstr, user, saldofloat, rd, data)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error" : err.Error(),
			})
			return 
		}

		ctx.JSON(200, gin.H{
			"operacion":"transferencia en proceso",
		})
	}
}