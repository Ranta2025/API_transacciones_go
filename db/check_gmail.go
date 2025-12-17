package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func Check_gmail(ctx context.Context, gmail string,rd *redis.Client, db *sql.DB) bool {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, err := rd.Get(c, fmt.Sprintf("gmailToID:%s", gmail)).Result()
	if err == redis.Nil {
		query := db.QueryRowContext(c, "SELECT id FROM transacciones.user WHERE gmail = ?", gmail)
		var id string
		err := query.Scan(&id)
		if err == sql.ErrNoRows {
			return false
		}
		rd.Set(c, fmt.Sprintf("gmailToID:%s", gmail), id, 2 * time.Minute)
		return true
	}else if err != nil {
		return false
	}
	_ = id
	return true
}