package controllers

import (
	"gosamples/gocommerce/models/products"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var htmlTemplates = template.Must(template.ParseGlob("resources/templates/*.html"))

// Index is
func Index(w http.ResponseWriter, r *http.Request) {
	products := products.FindAll()
	htmlTemplates.ExecuteTemplate(w, "Index", products)
}

// New is
func New(w http.ResponseWriter, r *http.Request) {
	htmlTemplates.ExecuteTemplate(w, "New", nil)
}

// Insert is
func Insert(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		name := req.FormValue("name")
		description := req.FormValue("description")
		amount := req.FormValue("amount")
		price := req.FormValue("price")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Impossible to convert price", err)
		}

		convertedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Impossible to convert amount", err)
		}

		products.Create(name, convertedPrice, convertedAmount, description)
	}
	http.Redirect(w, req, "/", 303)
}

// Delete is
func Delete(w http.ResponseWriter, req *http.Request) {
	// id := req.FormValue("id")
	id := req.URL.Query().Get("id")
	products.Delete(id)

	http.Redirect(w, req, "/", 303)
}

// Edit is
func Edit(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	product := products.FindByID(id)
	htmlTemplates.ExecuteTemplate(w, "Edit", product)
}

// Update is
func Update(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		name := req.FormValue("name")
		description := req.FormValue("description")
		amount := req.FormValue("amount")
		price := req.FormValue("price")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Impossible to convert price", err)
		}

		convertedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Impossible to convert amount", err)
		}
		id := req.FormValue("id")
		idConverted, _ := strconv.Atoi(id)

		// fmt.Println(idConverted, name, convertedPrice, convertedAmount, description)
		products.Update(idConverted, name, convertedPrice, convertedAmount, description)
	}
	http.Redirect(w, req, "/", 303)
}
