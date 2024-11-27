package user

import (
	"github.com/damasdev/boring-go/pkg/ctx"
	"github.com/gorilla/mux"
)

// Routes defines the routes for the user package.
func Routes(r *mux.Router) {
	r.HandleFunc("/login", ctx.Handler(HandleUserLogin)).Methods("GET")
	r.HandleFunc("/logout", ctx.Handler(HandleUserLogout)).Methods("GET")
}
