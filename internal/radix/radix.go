package radix

import "github.com/elmq0022/krillin/router"

type Node[T any] struct {
	prefix   string
	children []*Node[T]
	terminal map[string]T
}

type Radix[T any] struct {
	root *Node[T]
}

func New[T any]([]router.Route[T]) (*Radix[T], error) {
	return &Radix[T]{}, nil
}

func (r *Radix[T]) Lookup(method, path string) (T, bool) {
	return *new(T), true
}
