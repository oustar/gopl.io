package main

import (
	"net/http"
)

// Route is handers routes
type Route struct {
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is the slices of handler function
type Routes []Route

// NewMux create the new ServerMux
func NewMux() *http.ServeMux {

	mux := http.NewServeMux()
	for _, r := range routes {
		var handler http.Handler

		handler = r.HandlerFunc
		handler = Logger(handler)
		mux.Handle(r.Pattern, handler)
	}
	return mux
}

var routes = Routes{
	Route{
		"/todo/read",
		todoRead,
	},
	Route{
		"/todo/create",
		todoCreate,
	},
}
