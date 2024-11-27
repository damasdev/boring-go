package user

import (
	"net/http"

	"github.com/damasdev/boring-go/pkg/ctx"
)

// HandleUserLogin handles the login page.
func HandleUserLogin(c *ctx.Context) error {
	return c.Render(LoginPage())
}

// HandleUserLogout handles the logout page.
func HandleUserLogout(c *ctx.Context) error {
	return c.Redirect(http.StatusSeeOther, "/login")
}
