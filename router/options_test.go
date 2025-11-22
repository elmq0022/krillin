package router_test

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/elmq0022/kami/router"
	"github.com/elmq0022/kami/types"
)

func newTestMiddleware(val int) types.Middleware {
	return func(h types.Handler) types.Handler {
		return func(req *http.Request) (types.Response, error) {
			resp, err := h(req)
			bs, ok := resp.Body.([]int)
			if !ok {
				bs = []int{}
			}
			bs = append(bs, val)
			resp.Body = bs
			return resp, err
		}
	}
}

func testHandler(req *http.Request) (types.Response, error) {
	return types.Response{Status: 200, Body: []int{}}, nil
}

func TestWithMiddleware(t *testing.T) {
	spy := SpyAdapterRecord{}
	r, _ := router.New(
		NewSpyAdapter(&spy),
		router.WithMiddleware(newTestMiddleware(1)),
		router.WithMiddleware(newTestMiddleware(2)),
		router.WithMiddleware(newTestMiddleware(3)),
	)
	r.GET("/", testHandler)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(rec, req)

	if spy.Status != http.StatusOK {
		t.Fatalf("want %d got %d", http.StatusNotFound, spy.Status)
	}

	want := []int{3, 2, 1}
	got := spy.Body.([]int)
	if !slices.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestWithNotFound(t *testing.T) {
	testNotFound := func(r *http.Request) (types.Response, error) {
		return types.Response{
				Status: http.StatusNotFound,
				Body:   "test not found"},
			nil
	}
	spy := SpyAdapterRecord{}
	r, _ := router.New(NewSpyAdapter(&spy), router.WithNotFound(testNotFound))

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	r.ServeHTTP(rec, req)

	if spy.Status != http.StatusNotFound {
		t.Fatalf("want %d got %d", http.StatusNotFound, spy.Status)
	}

	if spy.Body != "test not found" {
		t.Fatalf("want %s, got %s", "test not found", spy.Body)
	}
}
