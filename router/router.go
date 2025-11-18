package router

import (
	"net/http"

	"github.com/elmq0022/krillin/internal/radix"
	"github.com/elmq0022/krillin/types"
)

type Router struct {
	routes  []types.Route
	adapter types.Adapter
	radix   *radix.Radix
}

func New(routes types.Routes, processor types.Adapter) *Router {
	radix, _ := radix.New(routes)

	return &Router{
		routes:  routes,
		adapter: processor,
		radix:   radix,
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h, _, ok := r.radix.Lookup(req.Method, req.URL.Path)
	if ok {
		r.adapter(w, req, h)
		return
	}
	http.NotFound(w, req)
}
