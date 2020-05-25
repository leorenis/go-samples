package db

import "database/sql"

// SetupDB is
func SetupDB() {
	db := OpenDBConnection()
	statement, _ := db.Prepare("create table if not exists products (id INTEGER PRIMARY KEY, name TEXT, price REAL, amount INTEGER, description TEXT)")
	statement.Exec()

	statement, _ = db.Prepare("DELETE FROM products")
	statement.Exec()

	statement, _ = db.Prepare("INSERT INTO products (name, price, amount, description) VALUES (?, ?, ?, ?)")
	statement.Exec("iPhone11 Pro", 1998., 1, "Plus")
	statement.Exec("Mackbook Pro", 2600., 2, "HD 512 SSD, Retina")
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
