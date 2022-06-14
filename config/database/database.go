package database

import (
	"database/sql"
)

// Model for collection of DB
type Database struct {
	MySQL *sql.DB
}
