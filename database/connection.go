package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "blogss"

	dateFormat     = "2006-01-02"
	datetimeFormat = "2006-01-02 15:04:05"
)

func Connect() (*sql.DB, error) {
	var connString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return sql.Open("postgres", connString)
}
