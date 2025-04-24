package users

import (
	"database/sql"
	"encoding/json"
	"gosamples/gousreadersqlite/db"
	"gosamples/gousreadersqlite/models/countries"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID      string `json:"id"`
	Name    string `json:"nome"`
	Country string `json:"pais"`
	Age     int    `json:"idade"`
	Score   int    `json:"score"`
	Active  bool   `json:"ativo"`
	Team    Team   `json:"equipe"`
	Logs    []Log  `json:"logs"`
}

type Team struct {
	Name     string    `json:"nome"`
	Leader   bool      `json:"lider"`
	Projects []Project `json:"projetos"`
}

type Project struct {
	Name string `json:"nome"`
	Done bool   `json:"concluido"`
}

type Log struct {
	Moment string `json:"data"`
	Action string `json:"acao"`
}

// Public route
func FindAllSuperUsers() []User {
	db := db.OpenDBConnection()
	rows, err := db.Query("select * from users where score >= 900 and active = 1 order by score desc")
	if err != nil {
		panic(err.Error())
	}

	// Referencias
	user := User{}
	users := []User{}

	for rows.Next() {
		var uuid, name, country string
		var active, age, score, id int

		err = rows.Scan(&id, &uuid, &name, &age, &active, &country, &score)
		if err != nil {
			panic(err.Error())
		}

		user.Score = score
		user.Country = country
		user.Age = age
		user.Name = name
		user.ID = uuid
		user.Active = active >= 1

		users = append(users, user)
	}
	defer db.Close()
	return users
}

func FindTopFiveCountries() []countries.Top {
	db := db.OpenDBConnection()
	rows, err := db.Query("select count(*) as total, u.country as country from users u group by country order by total desc limit 5")
	if err != nil {
		panic(err.Error())
	}

	top := countries.Top{}
	topFive := []countries.Top{}

	for rows.Next() {
		var total int
		var country string

		err = rows.Scan(&total, &country)
		if err != nil {
			panic(err.Error())
		}

		top.Total = total
		top.Country = country

		topFive = append(topFive, top)
	}
	defer db.Close()
	return topFive
}

// Public route
func Migrate() {
	db := db.OpenDBConnection()
	var count int
	row := db.QueryRow("SELECT Count(*) FROM users")

	switch err := row.Scan(&count); err {
	case nil:
		if count < 1 {
			log.Println("Iniciando carga...")
			for i := 1; i <= 99999; i++ {
				//for i := 1; i <= 9; i++ {
				doGetEndpoint(i)
			}
		}
	default:
		panic(err)
	}
	defer db.Close()
}

// Private func
func doGetEndpoint(id int) {
	response, err := http.Get("http://localhost:3000/" + strconv.Itoa(id))

	if err != nil {
		log.Println("Error: ", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode == 200 {

		var user User
		if err := json.Unmarshal(body, &user); err != nil {
			log.Fatal("Erro ao decodificar JSON: ", err)
		}
		// log.Println("User:", user)

		saveUser(&user)
	} else {
		log.Println("Error to get. ", response.StatusCode)
	}
}

// Private func
func saveUser(user *User) {
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
