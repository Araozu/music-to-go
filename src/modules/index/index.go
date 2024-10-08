package index

import (
	"acide/src/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Registers all the routes that this module (auth) provides
func SetupRoutes(g *echo.Group) {
	log.Print("Setting up the index module")

	// To include custom rendering logic:
	g.GET("/", indexPage)
}

func indexPage(c echo.Context) error {
	// If the required cookies are set, redirect to home
	_, err1 := c.Cookie("session-token")
	_, err2 := c.Cookie("navidrome-url")

	if err1 != nil || err2 != nil {
		return c.Redirect(http.StatusFound, "/auth/")
	}

	return utils.RenderTempl(c, http.StatusOK, IndexTempl())
}
