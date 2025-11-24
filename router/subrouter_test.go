package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elmq0022/kami/router"
)

func TestSubRouter(t *testing.T) {
	spy := SpyAdapterRecord{}
	r, err := router.New(NewSpyAdapter(&spy))
	if err != nil {
		t.Fatalf("%v", err)
	}

	api := r.Group("/api/v1/")
	wantStatus := 200
	wantBody := "bar"
	var wantErr error = nil

	api.GET("/foo", NewTestHandler(wantStatus, wantBody, wantErr))

	req, err := http.NewRequest(http.MethodGet, "/api/v1/foo", nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if spy.Status != wantStatus {
		t.Fatalf("want %d, got %d", wantStatus, spy.Status)
	}

	if spy.Body.(string) != wantBody {
		t.Fatalf("want %s, got %s", wantBody, spy.Body.(string))
	}

	if spy.Err != wantErr {
		t.Fatalf("want %v, got %v", &wantErr, spy.Err)
	}
}
