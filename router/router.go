package router

import (
	"net/http"
)

type Route[T any] struct {
	Method  string
	Path    string
	Handler T
}

type Adapter[T any] func(http.ResponseWriter, *http.Request, T)

type Router[T any] struct {
	routes    []Route[T]
	processor func(http.ResponseWriter, *http.Request, T)
}

func New[T any](routes []Route[T], processor Adapter[T]) *Router[T] {
	return &Router[T]{
		routes:    routes,
		processor: processor,
	}
}

func (r *Router[T]) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Method == req.Method && route.Path == req.URL.Path {
			r.processor(w, req, route.Handler)
			return
		}
	}
	http.NotFound(w, req)
}
