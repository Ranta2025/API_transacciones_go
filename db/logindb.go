package db

import (
	"context"
	"database/sql"
	"proyecto_transacciones/models"
	"proyecto_transacciones/utils"

	"github.com/redis/go-redis/v9"
)

func Logindb(c context.Context, data *sql.DB,rd *redis.Client, user models.Login) string {

	flag, userdb := Check_user(c, user.Username,rd , data)
	if !flag {
		return "usuario inexistente"
	}
	if userdb.Bloqueado == "si" {
		return "usuario bloqueado"
	}

	equals := utils.Compare_password(user.Password, userdb.Password)
	if !equals {
		Set_invalid(c, user.Username, data, rd)
		return "Contraseña incorrecta"
	}
	ResetError(c, user.Username, data)
	return ""
}