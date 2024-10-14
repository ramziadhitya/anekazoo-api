package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL() (*sql.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/anekazoo_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Cannot connect to MySQL database", err)
		return nil, err
	}

	log.Println("Connected to MySQL!")
	return db, nil
}
