package artist

import (
	"acide/src/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Setup(g *echo.Group) {
	log.Print("Setting up the artist module")
	g.Use(utils.Authed)

	// g.GET("/", allArtistPage)
	g.GET("/:id", artistPage)
}

func artistPage(c echo.Context) error {
	// token, server := utils.Credentials(c)
	isHtmxRequest := c.Request().Header.Get("HX-Request") == "true"
	// artistId := c.Param("id")

	// load artist info

	if isHtmxRequest {
		return utils.RenderTempl(c, http.StatusOK, artistTempl())
	} else {
		return utils.RenderTempl(c, http.StatusOK, artistTempl())
	}
}
