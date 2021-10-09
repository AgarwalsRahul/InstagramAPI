package router

import (
	"github.com/AgarwalsRahul/InstagramAPI/app"
	"github.com/julienschmidt/httprouter"
)

/*
Route is a struct for handling all routes
*/
type Route struct {
	Name            string
	Method          string
	Path            string
	HandlerFunction httprouter.Handle
}

// NewRouter is a helper function for creating new routes
func NewRouter(routes []Route) *httprouter.Router {

	if len(routes) <= 0 {
		return nil
	}

	router := httprouter.New()
	for _, route := range routes {
		handle := app.Logger(route.HandlerFunction)

		router.Handle(route.Method, route.Path, handle)
	}
	return router
}
