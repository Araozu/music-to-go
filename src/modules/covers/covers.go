package covers

import (
	"acide/src/utils"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Setup(g *echo.Group) {
	log.Print("Setting up the covers module")

	g.Use(utils.Authed)

	g.GET("/:id", getCover)
}

func getCover(c echo.Context) error {
	token, server := utils.Credentials(c)
	albumId := c.Param("id")

	coverBytes, err := loadCover(token, server, albumId)
	if err != nil {
		return err
	}

	c.Response().Header().Set("Cache-Control", "max-age=604800")

	return c.Blob(http.StatusOK, "image/png", coverBytes)
}
