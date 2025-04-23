package users

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
