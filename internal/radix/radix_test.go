package radix_test

import (
	"net/http"
	"testing"

	"github.com/elmq0022/krillin/internal/radix"
	"github.com/elmq0022/krillin/router"
)

func MakeTestHandler(value any) router.Handler {
	return func(req *http.Request) (int, any, error) {
		return 0, value, nil
	}
}

func ReadTestHandler(h router.Handler) any {
	fakeReq, _ := http.NewRequest(http.MethodGet, "", nil)
	_, value, _ := h(fakeReq)
	return value
}

func TestNewRadix(t *testing.T) {

	path := "/foo/bar/baz"
	method := http.MethodGet
	handler := MakeTestHandler(1)

	routes := router.Routes{
		{Path: path, Method: method, Handler: handler},
		{Path: "/foo/bar/baz2", Method: http.MethodPatch, Handler: MakeTestHandler(2)},
	}

	r, _ := radix.New(routes)

	h, _ := r.Lookup(method, path)
	if got := ReadTestHandler(h); got != 1 {
		t.Fatalf("want %d, got %d", 1, got)
	}

	h, _ = r.Lookup(http.MethodPatch, "/foo/bar/baz2")
	if got := ReadTestHandler(h); got != 2 {
		t.Fatalf("want %d, got %d", 2, got)
	}
}
