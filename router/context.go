package router

import "context"

type contextKey string

const paramsKey contextKey = "paramsKey"

// WithParams adds URL parameters to the request context.
// This is used internally by the router to store matched path parameters.
func WithParams(ctx context.Context, params map[string]string) context.Context {
	return context.WithValue(ctx, paramsKey, params)
}

// GetParams extracts URL parameters from the request context.
// Parameters come from route definitions like "/users/:id" where :id becomes a parameter.
// Returns an empty map if no parameters are present in the context.
func GetParams(ctx context.Context) map[string]string {
	if p, ok := ctx.Value(paramsKey).(map[string]string); ok {
		return p
	}
	return make(map[string]string)
}
