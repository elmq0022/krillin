package adapters

import (
	"net/http"
)

type Handler func(req *http.Request) (int, any, error)
type Adapter func(http.ResponseWriter, *http.Request, Handler)
