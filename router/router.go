package router

import "net/http"

type Handler func(w http.ResponseWriter, req *http.Request)

type Route struct {
	Method  string
	Path    string
	Handler Handler
}

type Router struct {
	routes []Route
}

func New(routes []Route) *Router {
	return &Router{
		routes: routes,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Method == req.Method && route.Path == req.URL.Path {
			route.Handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
