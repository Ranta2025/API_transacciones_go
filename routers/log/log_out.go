package log

import (
	"proyecto_transacciones/handler"
	"proyecto_transacciones/middlewares"
	"proyecto_transacciones/models"
	"proyecto_transacciones/utils"
)
func Log_out(app *models.App){
	sub := utils.GetLogout(app)
	sub.Srouter.Use(middlewares.Validation_param_into)
	sub.Srouter.Use(middlewares.Validation_param_len)
	sub.Srouter.Use(middlewares.Validation_password)
	sub.Srouter.POST("/", handler.Handler_logout(sub.Depends.Db, sub.Depends.Rd))
}