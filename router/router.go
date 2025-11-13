package router

import "net/http"

type Router struct {
}

func New() *Router {
	return &Router{}
}

func (r *Router) Get(path string) {

}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}
