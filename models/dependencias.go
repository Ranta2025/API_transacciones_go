package models

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type Dependencias struct {
	Db *sql.DB
	Rd *redis.Client
}