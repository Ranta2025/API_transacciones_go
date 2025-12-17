package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"proyecto_transacciones/models"
	"proyecto_transacciones/utils"
	"time"

	"github.com/redis/go-redis/v9"
)

func Set_invalid(ctx context.Context, username string, data *sql.DB, rd *redis.Client) {
	c,cancel := context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()

	_, err := data.ExecContext(c, "UPDATE transacciones.bloqueado SET cantidad_error_is = cantidad_error_is + 1", username)
	utils.Check_err(err)

	query := data.QueryRowContext(c, "SELECT cantidad_error_is FROM transacciones.bloqueado WHERE username = ?", username)
	var cantidad int
	query.Scan(&cantidad)

	if cantidad == 5{
		_,err := data.ExecContext(c, "UPDATE transacciones.user SET bloqueado = ? WHERE username = ?", "si", username)
		utils.Check_err(err)
		ResetError(ctx, username, data)
	}
	query2 := data.QueryRowContext(c, "SELECT password, bloqueado FROM transacciones.user WHERE username = ?", username)
	var user models.Login
	user.Username = username
	query2.Scan(&user.Password, &user.Bloqueado)
	userjson,_ := json.Marshal(user)
	extraer := fmt.Sprintf("userCheckUser:%s", username)
	rd.Set(c, extraer, userjson, 2 * time.Minute)
} 



func ResetError(c context.Context,username string, data *sql.DB){
	ctx, cancel := context.WithTimeout(c, 5 * time.Second)
	defer cancel()

	data.ExecContext(ctx, "UPDATE transacciones.bloqueado SET cantidad_error_is = 0 WHERE username = ? AND cantidad_error_is > 0", username)
}