package ctx

import (
	"net/http"

	"github.com/a-h/templ"
)

// HandlerFunc is a function that handles an HTTP request.
type HandlerFunc func(c *Context) error

// Context is a struct that represents the context of an HTTP request.
type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

// Render renders a view component.
func (c *Context) Render(view templ.Component) error {
	return view.Render(c.Request.Context(), c.Response)
}

// Redirect redirects the request to a new URL.
func (c *Context) Redirect(status int, url string) error {
	http.Redirect(c.Response, c.Request, url, status)
	return nil
}

// Handler is a function that wraps a handler function with a context.
func Handler(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &Context{
			Response: w,
			Request:  r,
		}

		if err := h(c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
