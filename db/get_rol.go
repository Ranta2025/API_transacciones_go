package db

import (
	"context"
	"database/sql"
	"time"
)

func GetRol(c context.Context, username string, data *sql.DB) string {
	ctx, cancel := context.WithTimeout(c, 5 * time.Second)
	defer cancel()

	query := data.QueryRowContext(ctx, "SELECT rol FROM transacciones.user WHERE username = ?", username)
	var rol string
	query.Scan(&rol)

	return rol
}