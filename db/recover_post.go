package db

import (
	"context"
	"database/sql"
	"proyecto_transacciones/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

func Recoverdb(c context.Context, username string, codigo string,rd *redis.Client, data *sql.DB) (bool, string){
	ctx, cancel := context.WithTimeout(c, 10 * time.Second)
	defer cancel()
	flag, user := Check_user(c, username, rd, data)
	if !flag {
		return false, "usuario inexistente"
	}

	if user.Bloqueado == "no" {
		return false, "usuario no bloqueado"
	}

	codigodb, err := rd.Get(ctx,"gmailRecoverCode:"+username).Result() 
	if err == redis.Nil {
		return false, "codigo expirado"
	}else if err != nil {
		return false, "error al buscar codigo"
	}

	if codigodb != codigo {
		return false, "codigo incorrecto"
	}
	_, err1 := data.ExecContext(ctx, "UPDATE transacciones.user SET bloqueado = ? WHERE username = ?", "no", username)
	utils.Check_err(err1)
	return true, "usuario desbloqueado"
}