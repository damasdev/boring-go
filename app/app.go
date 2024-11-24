package app

import (
	"net/http"
	"strings"

	"github.com/damasdev/boring-go/app/auth"
	"github.com/damasdev/boring-go/pkg/server"
)

// Boot the application.
func Boot(s *server.Server) {
	// Get the router from the server.
	r := s.Router()

	// Serve static files from the public directory.
	fileServer := http.FileServer(http.Dir("public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public", restrictedFileServer(fileServer)))

	// Load the routes from the auth package.
	auth.Routes(r)
}

// restrictedFileServer is a middleware that prevents directory listing.
func restrictedFileServer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
