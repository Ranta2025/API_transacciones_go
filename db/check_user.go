package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"proyecto_transacciones/models"
	"strconv"
	"time"
	"github.com/redis/go-redis/v9"
)

func Check_user(ctx context.Context, username string,rd *redis.Client, db *sql.DB) (bool, *models.Login ) {
	c, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()
	extraer := fmt.Sprintf("userCheckUser:%s", username)
	extraerUser, err := rd.Get(c, extraer).Result()
	if err == redis.Nil {
		query := db.QueryRowContext(c, "SELECT password, bloqueado FROM transacciones.user WHERE username = ?", username)
		var user models.Login
		user.Username = username
		err := query.Scan(&user.Password, &user.Bloqueado)
		if err == sql.ErrNoRows{
			return false, nil
		}
		data, _ := json.Marshal(user)
		rd.Set(c, extraer, data, 2 * time.Minute)
		return true, &user
	}
	var user models.Login
	json.Unmarshal([]byte(extraerUser), &user)  
	return true, &user
}


func Check_saldo(ctx *context.Context, username string, db *sql.DB, rd *redis.Client) float64 {
	c, cancel := context.WithTimeout(*ctx, 5 * time.Second)
	defer cancel()
	saldord,err := rd.Get(c, fmt.Sprintf("getSaldo:%s", username)).Result()
	if err == redis.Nil {
		query := db.QueryRowContext(c, "SELECT saldo FROM user WHERE username = ?", username)
		var saldo float64
		query.Scan(&saldo)
		rd.Set(c, fmt.Sprintf("getSaldo:%s", username), saldo, 5 * time.Minute)
		return saldo
	}
	saldoFlaot,_ := strconv.ParseFloat(saldord, 64)
	return saldoFlaot
}