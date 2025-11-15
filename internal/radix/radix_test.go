package radix_test

import (
	"net/http"
	"testing"

	"github.com/elmq0022/krillin/adapters"
	"github.com/elmq0022/krillin/internal/radix"
	"github.com/elmq0022/krillin/router"
)

func TestRadix(t *testing.T) {
	routes := []router.Route[adapters.Handler]{}

	radix, _ := radix.New[adapters.Handler](routes)

	url := "/url/path/to/resource"
	method := http.MethodGet
	req, _ := http.NewRequest(method, url, nil)

	_, ok := radix.Lookup(req.Method, req.URL.Path)
	if !ok {
		t.Fatal("Lookup was not OK")
	}
	// handler(req)
}
