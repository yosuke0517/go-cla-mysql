package db

import (
	"database/sql"
	"fmt"
	"os"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
