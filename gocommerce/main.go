package main

import (
	"gosamples/gocommerce/db"
	"gosamples/gocommerce/models/products"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

var htmlTemplates = template.Must(template.ParseGlob("resources/templates/*.html"))

func main() {
	// setup database
	db.SetupDB()
	// Up server
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// index is
func index(w http.ResponseWriter, r *http.Request) {
	db := db.OpenDBConnection()
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
	defer db.Close()
}
