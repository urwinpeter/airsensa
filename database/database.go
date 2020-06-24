package database

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	
}

func GetData(now time.Time, past time.Time) *sql.Rows {}
	db, err := sql.Open("sqlite3", "data.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(
		"SELECT category, name, price, datetime FROM data WHERE datetime BETWEEN ? AND ?", 
		past[:19],
		now[:19],
	)
	if err != nil {
		log.Fatal(err)
	}
	return rows
	
}   