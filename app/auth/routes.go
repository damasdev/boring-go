package auth

import (
	"github.com/damasdev/boring-go/pkg/ctx"
	"github.com/gorilla/mux"
)

// Routes defines the routes for the auth package.
func Routes(r *mux.Router) {
	r.HandleFunc("/", ctx.Handler(HandleIndex)).Methods("GET")
	r.HandleFunc("/login", ctx.Handler(HandleLogin)).Methods("GET")
	r.HandleFunc("/logout", ctx.Handler(HandleLogout)).Methods("GET")
}
