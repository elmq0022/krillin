// Package handlers provides default HTTP handlers for common scenarios such as 404 errors.
package handlers

import (
	"net/http"

	"github.com/elmq0022/kami/types"
)

type defaultNotFoundResponder struct {
	status int
	body   string
}

// Respond writes the plain text response to the ResponseWriter.
func (dnf *defaultNotFoundResponder) Respond(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(dnf.status)
	w.Write([]byte(dnf.body))
}

// DefaultNotFoundHandler is the default 404 handler used by the router.
// Returns a plain text "Not Found" response with HTTP 404 status.
func DefaultNotFoundHandler(r *http.Request) types.Responder {
	return &defaultNotFoundResponder{status: http.StatusNotFound, body: "Not Found"}
}
