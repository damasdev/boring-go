package auth

import (
	"net/http"

	"github.com/damasdev/boring-go/pkg/ctx"
)

// HandleIndex handles the index page.
func HandleIndex(c *ctx.Context) error {
	return c.Render(LandingPage())
}

// HandleLogin handles the login page.
func HandleLogin(c *ctx.Context) error {
	return c.Render(LoginPage())
}

// HandleLogout handles the logout page.
func HandleLogout(c *ctx.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}
