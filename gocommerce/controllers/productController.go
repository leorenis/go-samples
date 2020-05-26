package controllers

import (
	"gosamples/gocommerce/models/products"
	"html/template"
	"net/http"
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
