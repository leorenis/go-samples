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
	products := products.FindAll()

	htmlTemplates.ExecuteTemplate(w, "Index", products)
}
