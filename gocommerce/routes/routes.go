package routes

import (
	"gosamples/gocommerce/controllers"
	"net/http"
)

// Load is
func Load() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
}
