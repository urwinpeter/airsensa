package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/urwinpeter/airsensa/lifecycle"
)

func main() {
	db, err := sql.Open("sqlite3", "data.db")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Welcome")
	lc := lifecycle.NewLifecycle(db)
	lc.Start()
}
