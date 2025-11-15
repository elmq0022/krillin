package router_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elmq0022/krillin/adapters"
	"github.com/elmq0022/krillin/router"
)

func TestRouter_GetRoute(t *testing.T) {
	result := make(map[string]bool)
	result["ok"] = true
	want, _ := json.Marshal(result)

	handler := func(req *http.Request) (int, any, error) {
		return http.StatusOK, result, nil
	}

	// TODO: can I kill the generic here if I always
	// just use the handler type? Probably
	routes := []router.Route[adapters.Handler]{
		{
			http.MethodGet,
			"/",
			handler,
		},
	}

	r := router.New(routes, adapters.JsonAdapter)

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

	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if string(got) != string(want) {
		t.Fatalf("want %s, got %s", want, got)
	}
}
