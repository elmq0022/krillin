package radix

import (
	"fmt"
	"strings"

	"github.com/elmq0022/krillin/router"
)

type Node struct {
	prefix   string
	children []*Node
	terminal map[string]router.Handler
}

type Radix struct {
	root *Node
}

func New(routes []router.Route) (*Radix, error) {
	r := Radix{root: &Node{}}

	for _, route := range routes {
		if len(route.Path) == 0 || route.Path[0] != '/' {
			return nil, fmt.Errorf("path must start with '/'")
		}

		segments := strings.Split(route.Path, "/")[1:]
		r.addRoute(route, r.root, segments, 0)
	}

	return &r, nil
}

func (r *Radix) addRoute(route router.Route, node *Node, segments []string, pos int) {
	if pos >= len(segments) {
		if node.terminal == nil {
			node.terminal = make(map[string]router.Handler)
		}
		node.terminal[route.Method] = route.Handler
		return
	}

	seg := segments[pos]

	for _, child := range node.children {
		if child.prefix == seg {
			r.addRoute(route, child, segments, pos+1)
			return
		}
	}

	n := &Node{prefix: seg}
	node.children = append(node.children, n)
	r.addRoute(route, n, segments, pos+1)
}

func (r *Radix) Lookup(method, path string) (router.Handler, bool) {
	root := r.root
	segments := strings.Split(path, "/")
	if segments[0] == "" {
		segments = segments[1:]
	}
	return lookup(root, method, segments, 0)
}

func lookup(node *Node, method string, segments []string, pos int) (router.Handler, bool) {
	var zero router.Handler

	if node == nil {
		return zero, false
	}

	if pos >= len(segments) {
		handler, ok := node.terminal[method]
		return handler, ok
	}

	for _, child := range node.children {
		if segments[pos] == child.prefix {
			h, ok := lookup(child, method, segments, pos+1)
			return h, ok
		}
	}

	return zero, false
}
