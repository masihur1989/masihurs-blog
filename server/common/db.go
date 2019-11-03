package common

import (
	// need this for go/sql
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	// needed for go/sql
	_ "github.com/go-sql-driver/mysql"
)

// DB -
var DB *sql.DB
var err error

// connection pool details
const (
	DefaultMaxIdleConns    = 5
	DefaultMaxOpenConns    = 10
	DefaultConnMaxLifetime = time.Minute * 5
)

// GetDB should almost always be used to init your DB access.
// Go will keep track of connections, etc. for you, so no need to instantiate a DB object with each request -- that's overkill.
func GetDB() *sql.DB {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	DB, err = sql.Open("mysql", connString)
	if err != nil {
		log.Println("ERROR: failed to establish db connection!")
	}
	// TODO needs to load the Connection string from env
	DB.SetMaxIdleConns(DefaultMaxIdleConns)
	DB.SetMaxOpenConns(DefaultMaxOpenConns)
	DB.SetConnMaxLifetime(DefaultConnMaxLifetime)
	return DB
}

// CloseDB close DB when done
func CloseDB() {
	DB.Close()
}
