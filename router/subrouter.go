package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/elmq0022/kami/types"
)

// SubRouter provides scoped route registration under a common path prefix.
// Routes registered on a SubRouter are automatically prefixed with the group's path.
type SubRouter struct {
	r      *Router
	prefix string
}

// NewSubRouter creates a new SubRouter with the given prefix.
// The prefix must start with "/" and cannot end with a wildcard "*".
// Panics if the prefix is invalid.
func NewSubRouter(r *Router, prefix string) SubRouter {
	if len(prefix) < 1 {
		panic("prefix cannot be an empty string")
	}

	if prefix[0] != '/' {
		panic(fmt.Sprintf("prefix %s does not start with a '/'", prefix))
	}

	trimedPrefix := strings.TrimRight(prefix, "/")
	if trimedPrefix[len(trimedPrefix)-1] == '*' {
		panic(fmt.Sprintf("prefix %s cannot end in a wildcard '*'", prefix))
	}

	return SubRouter{
		r:      r,
		prefix: trimedPrefix,
	}
}

func (s *SubRouter) add(method, path string, handler types.Handler) {
	fullPath := s.prefix + "/" + strings.TrimLeft(path, "/")
	s.r.add(method, fullPath, handler)
}

// GET registers a handler for GET requests at the given path, prefixed with the SubRouter's prefix.
// Path can include parameters and wildcards.
// Panics if the route cannot be registered.
func (s *SubRouter) GET(path string, handler types.Handler) {
	s.add(http.MethodGet, path, handler)
}

// POST registers a handler for POST requests at the given path, prefixed with the SubRouter's prefix.
// Path can include parameters and wildcards.
// Panics if the route cannot be registered.
func (s *SubRouter) POST(path string, handler types.Handler) {
	s.add(http.MethodPost, path, handler)
}

// PUT registers a handler for PUT requests at the given path, prefixed with the SubRouter's prefix.
// Path can include parameters and wildcards.
// Panics if the route cannot be registered.
func (s *SubRouter) PUT(path string, handler types.Handler) {
	s.add(http.MethodPut, path, handler)
}

// PATCH registers a handler for PATCH requests at the given path, prefixed with the SubRouter's prefix.
// Path can include parameters and wildcards.
// Panics if the route cannot be registered.
func (s *SubRouter) PATCH(path string, handler types.Handler) {
	s.add(http.MethodPatch, path, handler)
}

// DELETE registers a handler for DELETE requests at the given path, prefixed with the SubRouter's prefix.
// Path can include parameters and wildcards.
// Panics if the route cannot be registered.
func (s *SubRouter) DELETE(path string, handler types.Handler) {
	s.add(http.MethodDelete, path, handler)
}

// HEAD registers a handler for HEAD requests at the given path, prefixed with the SubRouter's prefix.
// Path can include parameters and wildcards.
// Panics if the route cannot be registered.
func (s *SubRouter) HEAD(path string, handler types.Handler) {
	s.add(http.MethodHead, path, handler)
}

// OPTIONS registers a handler for OPTIONS requests at the given path, prefixed with the SubRouter's prefix.
// Path can include parameters and wildcards.
// Panics if the route cannot be registered.
func (s *SubRouter) OPTIONS(path string, handler types.Handler) {
	s.add(http.MethodOptions, path, handler)
}
