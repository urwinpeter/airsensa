package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/urwinpeter/airsensa/lifecycle"
)

func main() {
	db, dbAccessErr := sql.Open("sqlite3", "data.db")
	defer db.Close()
	if dbAccessErr != nil {
		log.Fatal(dbAccessErr)
	}
	fmt.Println("Welcome")
	lc := lifecycle.NewLifecycle(db)
	lc.Start()
}
