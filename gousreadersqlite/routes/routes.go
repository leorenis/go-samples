package routes

import (
	"gosamples/gousreadersqlite/controllers"
	"gosamples/gousreadersqlite/middlewares"
	"net/http"
)

// Load is
func Load() {
	http.Handle("/superusers", middlewares.DurationTiming(http.HandlerFunc(controllers.Index)))
	http.HandleFunc("/top-countries", controllers.TopCountries)
	http.HandleFunc("/team-insights", controllers.TeamInsights)

	http.HandleFunc("/migrate", controllers.Migrate)
}
