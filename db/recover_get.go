package db

import (
	"context"
	"database/sql"
	"fmt"
	"proyecto_transacciones/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

func Check_gmail_invalid(c context.Context, gmail string,rd *redis.Client, data *sql.DB) (bool, string) {
	ctx, cancel := context.WithTimeout(c,10*time.Second)
	defer cancel()

	bloqueado, err := rd.Get(ctx, fmt.Sprintf("gmailCheckBloqued:%s", gmail)).Result()
	if err == redis.Nil {
		query := data.QueryRowContext(ctx, "SELECT bloqueado from transacciones.user WHERE gmail = ?", gmail)
		var bloqueado string
		err := query.Scan(&bloqueado)
		if err == sql.ErrNoRows {
			return false,"gmail inexistente"
		}
		rd.Set(ctx, fmt.Sprintf("gmailCheckBloqued:%s", gmail), bloqueado, 2 * time.Minute)
	} else if err != nil {
		return false, "error al consultar redis"
	}
	if bloqueado == "no" {
		return false, "usuario no bloqueado"
	}
	var username string
	var err2 error
	username, err2 = rd.Get(ctx, fmt.Sprintf("gmailToUsername:%s", gmail)).Result()
	if err2 == redis.Nil {
		query := data.QueryRowContext(ctx, "SELECT username from transacciones.user WHERE gmail = ?", gmail)
		err := query.Scan(&username)
		if err == sql.ErrNoRows {
			return false,"gmail inexistente"
		}
		rd.Set(ctx, fmt.Sprintf("gmailToUsername:%s", gmail), username, 2 * time.Minute)
	} else if err != nil {
		return false, "error al consultar redis"
	}
	codigo := utils.Generate()
	rd.Set(ctx, "gmailRecoverCode:"+username , codigo, 15 * time.Minute)
	fmt.Print(codigo)
	
	return true, "Codigo enviado"
}