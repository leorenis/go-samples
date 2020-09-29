package db

import (
	"database/sql"
	"log"
)

// SetupDB is
func SetupDB() {
	db := OpenDBConnection()
	statement, _ := db.Prepare("create table if not exists products (id INTEGER PRIMARY KEY, name TEXT, price REAL, amount INTEGER, description TEXT)")
	statement.Exec()

	var count int
	row := db.QueryRow("SELECT Count(*) FROM products")

	switch err := row.Scan(&count); err {
	case sql.ErrNoRows:
		statement, _ = db.Prepare("INSERT INTO products (name, price, amount, description) VALUES (?, ?, ?, ?)")
		statement.Exec("iPhone11 Pro", 1998., 1, "Plus")
		statement.Exec("Mackbook Pro", 2800., 2, "HD 255 SSD, Retina")

	case nil:
		log.Println("Database already loaded with", count, "records.")
	default:
		panic(err)

	}

	defer db.Close()
}

// OpenDBConnection is
func OpenDBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "resources/database/gocommerce.db")
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
