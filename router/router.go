package router

import (
	"net/http"

	"github.com/elmq0022/kami/internal/radix"
	"github.com/elmq0022/kami/types"
)

type Router struct {
	routes  []types.Route
	adapter types.Adapter
	radix   *radix.Radix
}

func New(routes types.Routes, adapter types.Adapter) (*Router, error) {
	rdx, err := radix.New(routes)
	if err != nil {
		return nil, err
	}

	return &Router{
		routes:  routes,
		adapter: adapter,
		radix:   rdx,
	}, nil
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h, params, ok := r.radix.Lookup(req.Method, req.URL.Path)
	if !ok {
		http.NotFound(w, req)
		return
	}

	ctx := WithParams(req.Context(), params)
	req = req.WithContext(ctx)

	r.adapter(w, req, h)
}
