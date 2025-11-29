package router

import (
	"log"
	"net/http"
	"time"

	"github.com/elmq0022/kami/types"
)

// Option is a function that configures a Router during initialization.
type Option func(r *Router)

// WithMiddleware adds global middleware to the router.
// Middleware is applied to all routes in the order it is registered.
// Multiple calls to WithMiddleware will append middleware to the chain.
func WithMiddleware(mw types.Middleware) Option {
	return func(r *Router) {
		r.global = append(r.global, mw)
	}
}

// WithNotFound sets a custom handler for 404 Not Found responses.
// If not specified, a default "Not Found" handler is used.
func WithNotFound(h types.Handler) Option {
	return func(r *Router) {
		r.notFound = h
	}
}

// WithLogger configures request logging for the router.
// Logs each request with method, path, status code, and duration.
func WithLogger() Option {
	return func(r *Router) {
		loggingMiddleware := func(next types.Handler) types.Handler {
			return func(req *http.Request) types.Responder {
				start := time.Now()

				// Call the next handler
				responder := next(req)

				// Wrap the responder to capture the response
				return &loggingResponder{
					inner:  responder,
					method: req.Method,
					path:   req.URL.Path,
					start:  start,
				}
			}
		}
		r.global = append(r.global, loggingMiddleware)
	}
}

type loggingResponder struct {
	inner  types.Responder
	method string
	path   string
	start  time.Time
}

func (l *loggingResponder) Respond(w http.ResponseWriter, req *http.Request) {
	// Wrap the ResponseWriter to capture status code
	lw := &loggingWriter{ResponseWriter: w, statusCode: 200}

	// Call the inner responder
	l.inner.Respond(lw, req)

	// Log after response is written
	duration := time.Since(l.start)
	log.Printf("%s %s - %d (%v)", l.method, l.path, lw.statusCode, duration)
}

type loggingWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lw *loggingWriter) WriteHeader(code int) {
	lw.statusCode = code
	lw.ResponseWriter.WriteHeader(code)
}
