package db

import (
	"context"
	"database/sql"
	"proyecto_transacciones/models"
	"time"

	"github.com/redis/go-redis/v9"
)

func Logout(ctx context.Context, user models.UserLogout,rd *redis.Client, db *sql.DB) string {
	c,cancel := context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()
	if flag,_ := Check_user(c, user.Username,rd,  db); flag{
			return "usuario ya existente"
	}
	if Check_gmail(c, user.Gmail,rd, db) {
		return "gmail ya existente"
	}
	_,err1 := db.ExecContext(c,"INSERT INTO transacciones.user (username, gmail, password) VALUES (?,?,?)", user.Username, user.Gmail, user.Password)
	if err1 != nil {
		return "problema al crear cuenta"
	}
	
	return ""
}