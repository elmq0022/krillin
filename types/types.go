// Package types defines the core types used throughout the kami framework.
package types

import (
	"net/http"
)

// Responder represents any type that can write an HTTP response.
// Implementations should write appropriate headers, status codes, and body content
// to the ResponseWriter.
type Responder interface {
	Respond(w http.ResponseWriter, r *http.Request)
}

// Response is a generic response container that can hold any data type.
// Status is the HTTP status code, and Body is the data to be sent in the response.
type Response struct {
	Status int
	Body   any
}

// Middleware wraps a Handler to provide cross-cutting functionality such as
// logging, authentication, or request modification. Middleware can be chained
// and is applied in the order registered.
type Middleware func(h Handler) Handler

// Handler is the primary function signature for handling HTTP requests in kami.
// It receives a request and returns a Responder that knows how to write the response.
type Handler func(req *http.Request) Responder

// Routes is a collection of Route definitions.
type Routes []Route

// Route represents a single HTTP route mapping.
type Route struct {
	Method  string
	Path    string
	Handler Handler
}
