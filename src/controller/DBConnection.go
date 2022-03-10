package controller

import (
	"database/sql"
	"fmt"
	"log"
)
import "github.com/go-sql-driver/mysql"

var db *sql.DB

func Connect() *sql.DB {
	cfg := mysql.Config{
		User:                 "monty",
		Passwd:               "montypassword",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "golang",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(err)
	}

	fmt.Println("Info: \tConnected successfully")
	return db
}

func Disconnect() {
	err := db.Close()
	if err != nil {
		_ = fmt.Errorf("error while closing %v", err)
	}
	fmt.Println("Info: \tDisconnected successfully")
}
