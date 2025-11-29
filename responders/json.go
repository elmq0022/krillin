// Package responders provides implementations of types.Responder for common response types
// including JSON responses and static file serving.
package responders

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type jsonResponder struct {
	body   any
	status int
}

// JSONResponse creates a responder that serializes the given body to JSON.
// The status parameter sets the HTTP status code (e.g., http.StatusOK).
// If status is 0, defaults to 200 OK.
// Panics during Respond if the body cannot be marshaled to JSON.
func JSONResponse(body any, status int) *jsonResponder {
	return &jsonResponder{body: body, status: status}
}

// Respond writes the JSON response to the ResponseWriter.
// Sets Content-Type to "application/json" and marshals the body.
// Panics if marshaling fails, which will be caught by the router's panic recovery.
func (r *jsonResponder) Respond(w http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal(r.body)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal JSON response: %v", err))
	}

	w.Header().Set("Content-Type", "application/json")
	if r.status > 0 {
		w.WriteHeader(r.status)
	}
	w.Write(data)
}

type jsonErrorResponder struct {
	status int
	msg    string
}

// JSONErrorResponse creates a responder that returns a JSON error message.
// The response will have Content-Type "application/problem+json" per RFC 7807.
// The msg parameter becomes the "msg" field in the JSON response.
// Panics if the error message cannot be marshaled.
func JSONErrorResponse(msg string, status int) *jsonErrorResponder {
	return &jsonErrorResponder{msg: msg, status: status}
}

type jsonError struct {
	Msg string `json:"msg"`
}

// Respond writes the error response to the ResponseWriter.
// Sets Content-Type to "application/problem+json" and marshals the error.
// Panics if marshaling fails, which will be caught by the router's panic recovery.
func (e *jsonErrorResponder) Respond(w http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal(jsonError{Msg: e.msg})
	if err != nil {
		panic(fmt.Sprintf("failed to marshal JSON error response: %v", err))
	}

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(e.status)
	w.Write(data)
}
