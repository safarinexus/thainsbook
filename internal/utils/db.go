package utils

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	var db *sql.DB

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("MYSQL_USER")
	cfg.Passwd = os.Getenv("MYSQL_PASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = "localhost"
	cfg.DBName = os.Getenv("MYSQL_DATABASE")

	log.Println("Connecting to DB...")
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	log.Println("Successfully connected to MySQL")
	return db, nil
}
