package landing

import (
	"github.com/damasdev/boring-go/pkg/ctx"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/", ctx.Handler(HandleLanding)).Methods("GET")
}
