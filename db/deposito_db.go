package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func DepositarDB(c context.Context, username string, saldo float64, data *sql.DB, rd *redis.Client) (string, error) {
	ctx, cancel := context.WithTimeout(c, 5 * time.Second)
	defer cancel()

	dep,err := data.ExecContext(ctx, "UPDATE transacciones.user SET saldo = saldo + ? WHERE username = ?", saldo, username)
	cant, _ := dep.RowsAffected()
	if err != nil || cant == 0{
		return "", errors.New("Operacion interrumpida")
	}
	query := data.QueryRowContext(ctx, "SELECT saldo FROM transacciones.user WHERE username = ?", username)
	var saldoactual float64
	query.Scan(&saldoactual)
	rd.Set(c, fmt.Sprintf("getSaldo:%s", username), saldoactual, 5 * time.Minute)
	return "deposito realizado", nil
}