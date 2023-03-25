package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "testuser"
	password = "0987654321"
	dbname   = "firstbase"
)

func main() {

	log.Println("started")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Connection error:", err)
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Ping error:", err)
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS links (short_code VARCHAR NOT NULL, origin_url TEXT UNIQUE, user_id VARCHAR, is_deleted BOOLEAN DEFAULT FALSE);")

	if err != nil {
		log.Println("table create error:", err)
		panic(err)
	}
}
