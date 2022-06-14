package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username string = os.Getenv("DB_USER")
	password string = os.Getenv("DB_PASS")
	host     string = os.Getenv("DB_HOST")
	database string = os.Getenv("DB_NAME")
	port     string = os.Getenv("DB_PORT")
)

func MySQLConn(dbase *Database) error {
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?parseTime=true")
	if err != nil {
		return err
	}

	(*dbase).MySQL = db
	return nil
}
