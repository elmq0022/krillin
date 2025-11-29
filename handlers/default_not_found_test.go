package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elmq0022/kami/handlers"
)

func TestDefaultNotFoundHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/foo", nil)
	responder := handlers.DefaultNotFoundHandler(r)
	responder.Respond(rr, r)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("want %d, got %d", http.StatusNotFound, rr.Code)
	}

	if rr.Body.String() != "Not Found" {
		t.Fatalf("want %s, got %s", "Not Found", rr.Body.String())
	}
}
