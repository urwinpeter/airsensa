package storage

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ItemsDB struct {
	conn *sql.DB
}

func NewItemsDB(dbconn *sql.DB) *ItemsDB {
	return &ItemsDB{dbconn}
}

func (data *ItemsDB) GetData(now, past time.Time) []byte {
	rows, queryError := data.conn.Query(
		"SELECT category, name, price, datetime FROM items WHERE datetime BETWEEN ? AND ?",
		past,
		now,
	)
	defer rows.Close()
	if queryError != nil {
		log.Fatal(queryError)
	}
	var item Datum
	var items []Datum

	for rows.Next() {
		scanError := rows.Scan(
			&item.Category,
			&item.Name,
			&item.Price,
			&item.Datetime,
		)
		if scanError != nil {
			log.Fatal(scanError)
		}

		items = append(items, item)
	}
	rowsError := rows.Err()
	if rowsError != nil {
		log.Fatal(rowsError)
	}
	byte, err := json.Marshal(items)
	if err != nil {
		log.Fatal(err)
	}
	return byte
}
