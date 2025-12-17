package router

import (
	"proyecto_transacciones/handler"
	"proyecto_transacciones/middlewares"
	"proyecto_transacciones/models"
	"proyecto_transacciones/utils"
)
func Router(app *models.App){ 
	sub := utils.GetTransfer(app)
	sub.Srouter.Use(middlewares.Auth(sub.Depends.Db, sub.Depends.Rd))
	sub.Srouter.Use(middlewares.ValidationUser)
	sub.Srouter.GET("/saldo", handler.Handler_user_getSaldo(sub.Depends.Db, sub.Depends.Rd))
	sub.Srouter.POST("/depositar", middlewares.MiddlewaresSaldo, handler.HandlerDeposito(sub.Depends.Db, sub.Depends.Rd))
	sub.Srouter.POST("/transferir", middlewares.MiddlewaresTransfer, middlewares.MiddlewaresSaldo, handler.HandlerUserTransferSaldo(sub.Depends.Db, sub.Depends.Rd))
}