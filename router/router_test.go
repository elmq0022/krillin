package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elmq0022/krillin/router"
)

func TestRoutter_GetRoute(t *testing.T) {
	handler := func(req *http.Request) (int, any, error) {
		result := make(map[string]bool)
		result["ok"] = true
		return http.StatusOK, result, nil
	}
	routes := []router.Route{
		{
			http.MethodGet,
			"/",
			handler,
		},
	}
	r := router.New(routes)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", res.StatusCode)
	}
	if got := res.Header.Get("Content-Type"); got != "application/json" {
		t.Fatalf("unexpected content-type: %q", got)
	}
}
