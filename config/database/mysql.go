package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "root"
	password string = ""
	host     string = "localhost"
	database string = "mezink"
	port     string = "3306"
)

func MySQLConn() (*Database, error) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database)
	if err != nil {
		return nil, err
	}
	return &Database{MySQL: db}, nil
}
