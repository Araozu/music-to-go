package auth

import (
	"acide/src/utils"
	"net/http"

	"github.com/labstack/echo"
)

// Renders the login form
func login(c echo.Context) error {

	return utils.RenderTempl(c, http.StatusOK, LoginTempl())
}
