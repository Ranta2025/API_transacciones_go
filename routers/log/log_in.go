package log

import (
	"proyecto_transacciones/handler"
	"proyecto_transacciones/middlewares"
	"proyecto_transacciones/models"
	"proyecto_transacciones/utils"
)
func Log_in(app *models.App){
	sub := utils.GetLogin(app)
	sub.Srouter.Use(middlewares.Validation_Login)
	sub.Srouter.POST("/", handler.Handler_login(sub.Depends.Db, sub.Depends.Rd))
}