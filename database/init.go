package database

import (
	"database/sql"
	"fmt"
	
	"log"
)

func InitDB() (*sql.DB, error) {
	fmt.Println("Server is running on port 8080")
	db, err := sql.Open("sqlite3", "./goexpert.db")
	if err != nil {
		log.Fatalln(err)
	}

	sqlStmt := `
	create table if not exists
		cotacao (id integer not null primary key,
			code text,
			codein text,
			name text,
			high text,
			low text,
			varBid text,
			pctChange text,
			bid text,
			ask text,
			timestamp text,
			create_date text);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	return db, nil
}
