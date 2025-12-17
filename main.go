package main

import (
	"context"
	"log"
	"os"
	"proyecto_transacciones/db"
	"proyecto_transacciones/models"
	"proyecto_transacciones/routers/routes"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	app := GetApp()
	defer app.Depend.Db.Close()
	app.Getbienvenida()
	routes.Conect(app)
	app.Run()
}

func GetApp() *models.App {
	app := gin.Default()

	db := db.Get_Db()
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	addr := os.Getenv("REDIS_HOST")
	if addr == "" {
		addr = "localhost:6379"
	}

	rd := redis.NewClient(&redis.Options{

		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err := rd.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	instance := models.App{Router: app, Depend: models.Dependencias{Db: db, Rd: rd}}
	return &instance
}
