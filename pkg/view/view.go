package view

import (
	"fmt"
	"time"
)

// version is the version of the JS and CSS files.
var version = time.Now().Unix()

// JS function to return the path of the JS file.
func JS(name string) string {
	return fmt.Sprintf("/public/js/%s?v=%d", name, version)
}

// CSS returns the path to the CSS file.
func CSS(name string) string {
	return fmt.Sprintf("/public/css/%s?v=%d", name, version)
}
