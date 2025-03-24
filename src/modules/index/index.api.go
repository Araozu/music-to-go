package index

import (
	"acide/src/utils"

	"github.com/labstack/echo/v4"
)

func randomAlbumsApi(c echo.Context) error {
	sessionToken, navidromeUrl := utils.Credentials(c)

	albums, err := getRandomAlbums(sessionToken, navidromeUrl, 10)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, albums)
}
