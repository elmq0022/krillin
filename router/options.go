package router

type Option func(r *Router)

func WithMiddleware() Option {
	return func(r *Router) {}
}

func WithNotFound() Option {
	return func(r *Router) {}
}

func WithLogger() Option {
	return func(r *Router) {}
}
