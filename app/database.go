package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewDB() *sql.DB {
	connect := GetEnvVariable("DB_USERNAME") +
		":" + "@tcp(" + GetEnvVariable("DB_HOST") +
		":" + GetEnvVariable("DB_PORT") +
		")/" + GetEnvVariable("DB_NAME") +
		"?parseTime=True"

	db, err := sql.Open("mysql", connect)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
