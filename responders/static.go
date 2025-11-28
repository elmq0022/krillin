package responders

import (
	"io/fs"
	"net/http"
	"strings"
)

type staticDirectoryResponder struct {
	FS      fs.FS
	Prefix  string
	handler http.Handler
}

func NewStaticDirResponder(f fs.FS, prefix string) *staticDirectoryResponder {
	fsHandler := http.StripPrefix(prefix, http.FileServer(http.FS(f)))

	return &staticDirectoryResponder{
		FS:      f,
		Prefix:  prefix,
		handler: fsHandler,
	}
}

func (r *staticDirectoryResponder) Respond(w http.ResponseWriter, req *http.Request) {
	trimmed := strings.TrimPrefix(req.URL.Path, r.Prefix)

	// If the URL path does not end with "/" and is a directory (or empty), redirect
	if !strings.HasSuffix(req.URL.Path, "/") {
		// Empty path is the root of FS
		if trimmed == "" {
			http.Redirect(w, req, req.URL.Path+"/", http.StatusMovedPermanently)
			return
		}

		// Otherwise, check FS
		if dir, err := r.FS.Open(trimmed); err == nil {
			if info, err := dir.Stat(); err == nil && info.IsDir() {
				http.Redirect(w, req, req.URL.Path+"/", http.StatusMovedPermanently)
				return
			}
		}
	}

	r.handler.ServeHTTP(w, req)
}
