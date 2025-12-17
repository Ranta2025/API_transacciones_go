package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func Get_Db() *sql.DB{
	dsn := "root:Rafaeltanda@123@tcp(localhost:3306)/transacciones"

	db, err := sql.Open("mysql", dsn)
	
	if err != nil{
		log.Fatal(err)
	}	

	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}
	
	return db
}
