package routes

import (
	"gosamples/gocommerce/controllers"
	"net/http"
)

// Load is
func Load() {
	http.HandleFunc("/", controllers.Index)
}
