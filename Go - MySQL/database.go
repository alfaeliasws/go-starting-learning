package go_mysql_api

import (
	"database/sql"
	"time"

	// "fml"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_mysql_database")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db

	//db usage

}
