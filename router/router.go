package router

import (
	"encoding/json"
	"net/http"
)

type Handler func(req *http.Request) (int, any, error)

func Handler2Json(w http.ResponseWriter, req *http.Request, handler Handler) {
	status, result, _ := handler(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	data, _ := json.Marshal(result)
	w.Write([]byte(data))
}

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
			Handler2Json(w, req, route.Handler)
			return
		}
	}
	http.NotFound(w, req)
}
