package service

import (
	"log"
	"time"

	"github.com/urwinpeter/airsensa/database"
	"github.com/urwinpeter/airsensa/storage"
)

//var end_time time.Time

// Load 10 days of data into cache
func Load() {
	now := time.Now()
	rows := database.GetData(
		now,
		now.Add(time.Hour*24*10*-1),
	)

	var (
		category string
		name     string
		price    float32
		datetime string
	)

	for rows.Next() {
		err := rows.Scan(
			&category,
			&name,
			&price,
			&datetime,
		)
		if err != nil {
			log.Fatal(err)
		}
		storage.Load(category, name, price, datetime)
	}

}
