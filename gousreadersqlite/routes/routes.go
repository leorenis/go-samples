package routes

import (
	"gosamples/gousreadersqlite/controllers"
	"net/http"
)

// Load is
func Load() {
	http.HandleFunc("/superusers", controllers.Index)
	http.HandleFunc("/top-countries", controllers.TopCountries)
	http.HandleFunc("/team-insights", controllers.TeamInsights)

	http.HandleFunc("/migrate", controllers.Migrate)
}
