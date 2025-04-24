package routes

import (
	"gosamples/gousreadersqlite/controllers"
	"net/http"
)

// Load is
func Load() {
	http.HandleFunc("/superusers", controllers.Index)

	http.HandleFunc("/migrate", controllers.Migrate)
}
