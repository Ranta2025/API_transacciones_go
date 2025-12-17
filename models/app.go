package models

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
    Depend Dependencias
}

func (a *App) Getbienvenida(){
	a.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Bienvenido": "Transacciones SA",
		})
	})
}

func (a *App) Run(){
	a.Router.Run(":8080")
}

