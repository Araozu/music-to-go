package index

import (
	"acide/src/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func SetupRoutes(g *echo.Group) {
	log.Print("Setting up the index module")

	g.Use(utils.Authed)

	// To include custom rendering logic:
	g.GET("/", indexPage)
}

func indexPage(c echo.Context) error {
	refreshQuery := c.QueryParam("refresh")

	sessionToken, navidromeUrl := utils.Credentials(c)
	albums, err := getRandomAlbums(sessionToken, navidromeUrl, 10)
	if err != nil {
		return c.HTML(http.StatusBadRequest, fmt.Sprintf("%s", err))
	}

	if refreshQuery == "true" {
		return utils.RenderTempl(c, http.StatusOK, RandomAlbumsFragment(albums))
	}

	return utils.RenderTempl(c, http.StatusOK, IndexTempl(albums))
}
