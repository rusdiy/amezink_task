package database

import (
	"database/sql"
)

type Database struct {
	MySQL *sql.DB
}