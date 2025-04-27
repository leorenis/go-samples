package routes

import (
	"gosamples/gousreadersqlite/controllers"
	"gosamples/gousreadersqlite/middlewares"
	"net/http"
)

// Load is
func Load() {
	http.Handle("/superusers", middlewares.DurationTiming(http.HandlerFunc(controllers.Index)))
	http.Handle("/top-countries", middlewares.DurationTiming(http.HandlerFunc(controllers.TopCountries)))
	http.Handle("/team-insights", middlewares.DurationTiming(http.HandlerFunc(controllers.TeamInsights)))

	http.HandleFunc("/migrate", controllers.Migrate)
}
