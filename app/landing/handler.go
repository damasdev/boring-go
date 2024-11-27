package landing

import "github.com/damasdev/boring-go/pkg/ctx"

// HandleIndex handles the index page.
func HandleLanding(c *ctx.Context) error {
	return c.Render(LandingPage())
}
