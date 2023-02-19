package routes

import "net/http"

// Route is a struct that represents a route
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	Autheticated bool
}
