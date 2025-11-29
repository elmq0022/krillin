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

// NewStaticDirResponder creates a responder that serves static files from the given filesystem.
// The prefix is the URL path prefix that will be stripped before looking up files in the FS.
// For example, with prefix "/static" and FS containing "index.html",
// a request to "/static/index.html" will serve the file.
// Delegates to http.FileServer for actual file serving.
func NewStaticDirResponder(f fs.FS, prefix string) *staticDirectoryResponder {
	fsHandler := http.StripPrefix(prefix, http.FileServer(http.FS(f)))

	return &staticDirectoryResponder{
		FS:      f,
		Prefix:  prefix,
		handler: fsHandler,
	}
}

// Respond serves static files from the configured filesystem.
// Automatically redirects directory requests to include a trailing slash.
// For example, "/static/dir" redirects to "/static/dir/" with a 301 status.
// Delegates to the underlying http.FileServer for actual file serving and security.
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
