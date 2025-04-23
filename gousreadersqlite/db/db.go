package db

import (
	"database/sql"
	"encoding/json"
	"gosamples/gousreadersqlite/models/users"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
		log.Println("Iniciando carga...")
		for i := 1; i <= 99999; i++ {
			//for i := 1; i <= 9; i++ {
			doGetEndpoint(i)
		}

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

func doGetEndpoint(id int) {
	response, err := http.Get("http://localhost:3000/" + strconv.Itoa(id))

	if err != nil {
		log.Println("Error: ", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode == 200 {

		var user users.User
		if err := json.Unmarshal(body, &user); err != nil {
			log.Fatal("Erro ao decodificar JSON: ", err)
		}
		// log.Println("User:", user)

		saveUser(&user)
	} else {
		log.Println("Error to get. ", response.StatusCode)
	}
}

func saveUser(user *users.User) {
	db, err := sql.Open("sqlite3", "resources/database/users.db")
	statement, _ := db.Prepare("INSERT INTO users (uuid, name, age, active, country, score) VALUES (?, ?, ?, ?, ?, ?)")
	rs, _ := statement.Exec(user.ID, user.Name, user.Age, bool2Int(user.Active), user.Country, user.Score)

	userid, errUsrId := rs.LastInsertId()
	if errUsrId == nil {
		statement, _ = db.Prepare("INSERT INTO teams (name, leader, userid) VALUES (?, ?, ?)")
		rs, _ = statement.Exec(user.Team.Name, bool2Int(user.Team.Leader), userid)

		statement, _ = db.Prepare("INSERT INTO logs (moment, action, userid) VALUES (?, ?, ?)")
		for _, log2Save := range user.Logs {
			rs, _ = statement.Exec(log2Save.Moment, log2Save.Action, userid)
		}

		teamid, errTeamId := rs.LastInsertId()
		if errTeamId == nil {
			statement, _ = db.Prepare("INSERT INTO projects (name, done, teamid) VALUES (?, ?, ?)")
			for _, project := range user.Team.Projects {
				rs, _ = statement.Exec(project.Name, bool2Int(project.Done), teamid)
			}
		}
	}

	if err != nil {
		log.Println("Error: ", err)
	}
	defer db.Close()
}

func bool2Int(b bool) int {
	if b {
		return 1
	}
	return 0
}
