package main

import (
	"net/http"
	"text/template"
)

var htmlTemplates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

// index is
func index(w http.ResponseWriter, r *http.Request) {
	htmlTemplates.ExecuteTemplate(w, "Index", nil)
}
