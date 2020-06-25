package storage

import (
	"database/sql"
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

func (data *ItemsDB) GetData(now, past time.Time) []Datum {
	rows, err := data.conn.Query(
		"SELECT category, name, price, datetime FROM items WHERE datetime BETWEEN ? AND ?",
		past,
		now,
	)
	if err != nil {
		log.Fatal(err)
	}
	var item Datum
	var items []Datum

	for rows.Next() {
		err := rows.Scan(
			&item.Category,
			&item.Name,
			&item.price,
			&item.datetime,
		)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	return items
}
