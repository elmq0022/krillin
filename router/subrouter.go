package router

import (
	"net/http"
	"strings"

	"github.com/elmq0022/kami/types"
)

type SubRouter struct {
	r      *Router
	prefix string
}

func NewSubRouter(r *Router, prefix string) SubRouter {
	return SubRouter{
		r:      r,
		prefix: prefix,
	}
}

func (s *SubRouter) add(method, path string, handler types.Handler) {
	path = strings.Join([]string{s.prefix, path}, "/")
	s.r.add(method, path, handler)
}

func (s *SubRouter) GET(path string, handler types.Handler) {
	s.add(http.MethodGet, path, handler)
}

func (s *SubRouter) POST(path string, handler types.Handler) {
	s.add(http.MethodPost, path, handler)
}

func (s *SubRouter) PUT(path string, handler types.Handler) {
	s.add(http.MethodPut, path, handler)
}

func (s *SubRouter) PATCH(path string, handler types.Handler) {
	s.add(http.MethodPatch, path, handler)
}

func (s *SubRouter) DELETE(path string, handler types.Handler) {
	s.add(http.MethodDelete, path, handler)
}

func (s *SubRouter) HEAD(path string, handler types.Handler) {
	s.add(http.MethodHead, path, handler)
}

func (s *SubRouter) OPTIONS(path string, handler types.Handler) {
	s.add(http.MethodOptions, path, handler)
}
