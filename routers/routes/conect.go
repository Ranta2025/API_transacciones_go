package routes

import (
	"proyecto_transacciones/models"
	"proyecto_transacciones/routers/log"
	"proyecto_transacciones/routers/router"
)

func Conect(rt *models.App){
	log.Log_out(rt)
	log.Log_in(rt)
	router.Router(rt)
	router.Router_recover(rt)
}