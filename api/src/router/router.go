package router

import "github.com/gorilla/mux"

// Generate returns a router with all the routes registered
func Generate() *mux.Router {
	return mux.NewRouter()
}
