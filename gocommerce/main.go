package main

import (
	"gosamples/gocommerce/products"
	"net/http"
	"text/template"
)

var htmlTemplates = template.Must(template.ParseGlob("resources/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// index is
func index(w http.ResponseWriter, r *http.Request) {
	products := []products.Product{
		{"Mackbook Pro", 2600., 1, "HD 512 SSD, Retina"},
		{"iPhone11 Pro", 1998., 1, "Plus"},
	}
	htmlTemplates.ExecuteTemplate(w, "Index", products)
}
