package utils

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo"
)

// Middleware that allows only requests with the `session-token` and `navidrome-url` cookies set
func Authed(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("session-token")
		if err != nil {
			c.Redirect(http.StatusFound, "/auth/")
			return nil
		}
		_, err = c.Cookie("navidrome-url")
		if err != nil {
			c.Redirect(http.StatusFound, "/auth/")
			return nil
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}

// Renders a template and sends it with a custom http status code
func RenderTempl(c echo.Context, status int, cmp templ.Component) error {
	var buff bytes.Buffer
	if err := cmp.Render(c.Request().Context(), &buff); err != nil {
		log.Print(err)
		return errors.New("Error rendering templ component")
	}

	return c.HTML(status, buff.String())
}

// Returns the `sessionToken` and `navidromeUrl` cookies.
// This function must be called by a route protected by the Auth
// middleware, otherwise it will panic
func Credentials(c echo.Context) (string, string) {
	sessionToken, err := c.Cookie("session-token")
	if err != nil {
		panic("Error getting credentials from cookie: session-token was not set")
	}

	navidromeUrl, err := c.Cookie("navidrome-url")
	if err != nil {
		panic("Error getting credentials from cookie: navidrome-url was not set")
	}

	return sessionToken.Value, navidromeUrl.Value
}

func EscapeSingle(s string) string {
	return strings.ReplaceAll(s, "'", "\\'")
}
