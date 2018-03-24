package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route definiation
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes Array of Route
type Routes []Route

// NewRouter create needed Routers
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		indexHandler,
	},
	Route{
		"ComingFilm",
		"GET",
		"/",
		comingHandler,
	},
	Route{
		"FilmDetail",
		"GET",
		"/film/{filmId}",
		filmHandler,
	},
}
