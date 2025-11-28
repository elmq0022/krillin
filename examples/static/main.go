package main

import (
	"os"
	"path/filepath"

	"github.com/elmq0022/kami/router"
)

func main() {
	r, err := router.New()
	if err != nil {
		panic(err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		panic("could not find home dir")
	}
	base := filepath.Join(home, "web")

	f := os.DirFS(base)

	// serve static files by passing an fs.FS base directory
	// and the route base rout to r.ServeStatic
	r.ServeStatic(f, "/web/")

	// run the app
	r.Run(":8080")
}
