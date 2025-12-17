package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func TranderSaldo(c context.Context, username string, user string, saldo float64, rd *redis.Client, data *sql.DB) error {
	ctx, cancel := context.WithTimeout(c, 15 * time.Second)
	defer cancel()

	if exist,_:= Check_user(ctx, user, rd, data); !exist {
		return errors.New("usuario receptor inexistente")
	} 

	saldoDB := data.QueryRowContext(ctx, "SELECT saldo FROM transacciones.user WHERE username = ?", username)
	var saldoExtraido float64
	saldoDB.Scan(saldoExtraido)

	if saldoExtraido > saldo {
		return errors.New("Saldo insuficiente")
	}

	tx, err := data.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return errors.New("error al iniciar transaccion")
	}

	transaccion, err2 := tx.ExecContext(ctx, "UPDATE transacciones.user SET saldo = saldo - ? WHERE username = ?", saldo, username)
	cant, _ := transaccion.RowsAffected()
	if err2 != nil || cant == 0{
		tx.Rollback()
		return errors.New("error en transaccion")
	}

	deposito, err3 := tx.ExecContext(ctx, "UPDATE transacciones.user SET saldo = saldo + ? WHERE username = ?", saldo, user)
	cant2,_ := deposito.RowsAffected()

	if err3 != nil || cant2 == 0 {
		tx.Rollback()
		return errors.New("error en transaccion")
	}
	tx.Commit()

	_, err4 := rd.Get(ctx, fmt.Sprintf("getSaldo:%s", username)).Result()
	if err4 != redis.Nil {
		querySaldo1 := data.QueryRowContext(ctx, "SELECT saldo FROM transacciones.user WHERE username = ?", username)
		var saldoUser1 float64
		querySaldo1.Scan(&saldoUser1)
		rd.Set(ctx, fmt.Sprintf("getSaldo:%s", username), saldoUser1, 2 * time.Minute)
	}
	
	_, err5 := rd.Get(ctx, fmt.Sprintf("getSaldo:%s", user)).Result()
	if err5 != redis.Nil || err5 != nil{
		querySaldo2 := data.QueryRowContext(ctx, "SELECT saldo FROM transacciones.user WHERE username = ?", username)
		var saldoUser2 float64
		querySaldo2.Scan(&saldoUser2)
		rd.Set(ctx, fmt.Sprintf("getSaldo:%s", user), saldoUser2, 2 * time.Minute)
	} 

	return nil
}