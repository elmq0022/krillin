package types

import (
	"net/http"
)

type Handler func(req *http.Request) (int, any, error)
type Routes []Route
type Adapter func(http.ResponseWriter, *http.Request, Handler)

type Route struct {
	Method  string
	Path    string
	Handler Handler
}
