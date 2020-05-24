package main

import (
	"database/sql"
	"gosamples/gocommerce/products"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

var htmlTemplates = template.Must(template.ParseGlob("resources/templates/*.html"))

func main() {
	// setup database
	setupDB()
	// Up server
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// index is
func index(w http.ResponseWriter, r *http.Request) {
	db := openDBConnection()
	rows, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error)
	}
	product := products.Product{}
	products := []products.Product{}

	for rows.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &price, &amount, &description)
		if err != nil {
			panic(err.Error())
		}

		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		products = append(products, product)
	}

	htmlTemplates.ExecuteTemplate(w, "Index", products)
	db.Close()
}

func setupDB() {
	db := openDBConnection()
	statement, _ := db.Prepare("create table if not exists products (id INTEGER PRIMARY KEY, name TEXT, price REAL, amount INTEGER, description TEXT)")
	statement.Exec()

	statement, _ = db.Prepare("DELETE FROM products")
	statement.Exec()

	statement, _ = db.Prepare("INSERT INTO products (name, price, amount, description) VALUES (?, ?, ?, ?)")
	statement.Exec("iPhone11 Pro", 1998., 1, "Plus")
	statement.Exec("Mackbook Pro", 2600., 2, "HD 512 SSD, Retina")
	defer db.Close()
}

func openDBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "resources/database/gocommerce.db")
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}
