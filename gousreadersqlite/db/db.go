package db

import (
	"database/sql"
	"log"
)

// SetupDB is
func SetupDB() {
	db := OpenDBConnection()
	statement, _ := db.Prepare("create table if not exists users (id INTEGER PRIMARY KEY, uuid TEXT, name TEXT, age INTEGER, active INTEGER, country TEXT, score INTEGER)")
	statement.Exec()
	log.Println("table users created")

	statement, _ = db.Prepare("create table if not exists teams (id INTEGER PRIMARY KEY, name TEXT, leader INTEGER, userid INTEGER, FOREIGN KEY(userid) REFERENCES users(id))")
	statement.Exec()
	log.Println("table teams created")

	statement, _ = db.Prepare("create table if not exists projects (id INTEGER PRIMARY KEY, name TEXT, done INTEGER, teamid INTEGER, FOREIGN KEY(teamid) REFERENCES teams(id))")
	statement.Exec()
	log.Println("table projects created")

	statement, _ = db.Prepare("create table if not exists logs (id INTEGER PRIMARY KEY, moment TEXT, action TEXT, userid INTEGER, FOREIGN KEY(userid) REFERENCES users(id))")
	statement.Exec()
	log.Println("table logs created")

	var count int
	row := db.QueryRow("SELECT Count(*) FROM users")

	switch err := row.Scan(&count); err {
	case sql.ErrNoRows:
		statement, _ = db.Prepare("INSERT INTO users (uuid, name, age, active, country) VALUES (?, ?, ?, ?, ?, ?)")
		statement.Exec("abcef-iabcbc", "Joao", 20, 1, "Brazil", 1)

	case nil:
		log.Println("Database already loaded with", count, "records.")
	default:
		panic(err)
	}

	defer db.Close()
}

// OpenDBConnection is
func OpenDBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "resources/database/users.db")
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
