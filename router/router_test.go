package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elmq0022/krillin/router"
)

func TestRoutter_GetRoute(t *testing.T) {
	routes := []router.Route{
		{
			http.MethodGet,
			"/",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"ok":true}`))
			},
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
