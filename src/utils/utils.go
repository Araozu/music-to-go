package utils

import (
	"bytes"
	"errors"
	"log"

	"github.com/a-h/templ"
	"github.com/labstack/echo"
)

// Renders a template and sends it with a custom http status code
func RenderTempl(c echo.Context, status int, cmp templ.Component) error {
	var buff bytes.Buffer
	if err := cmp.Render(c.Request().Context(), &buff); err != nil {
		log.Print(err)
		return errors.New("Error rendering templ component")
	}

	return c.HTML(status, buff.String())
}
