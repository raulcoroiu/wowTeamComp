package routes

import "net/http"

type RouteEntry struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

type Router struct {
	routes []RouteEntry
}

func (rtr *Router) Route(method, path string, handlerFunc http.HandlerFunc) {
	e := RouteEntry{
		Method:  method,
		Path:    path,
		Handler: handlerFunc,
	}

	rtr.routes = append(rtr.routes, e)
}
