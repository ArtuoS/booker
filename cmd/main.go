package main

import (
	"database/sql"

	"github.com/ArtuoS/booker-api/internal/service"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../internal/database/booker.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	csvReader := service.NewCsvReader("../assets/files/authors.csv", db)
	csvReader.Start()
}
