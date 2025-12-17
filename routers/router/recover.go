package router

import (
	"proyecto_transacciones/handler"
	"proyecto_transacciones/middlewares"
	"proyecto_transacciones/models"
)

func Router_recover(app *models.App) {
	sub := models.Sublog{
		Srouter: app.Router.Group("/recover"),
		Depends: &app.Depend,
	}
	sub.Srouter.GET("/", middlewares.Validation_recover, handler.Handler_Recover(sub.Depends.Db,sub.Depends.Rd))
	sub.Srouter.POST("/", middlewares.Validation_recover_post, handler.Recover_post(sub.Depends.Db, sub.Depends.Rd))
}