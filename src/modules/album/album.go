package album

import (
	"acide/src/modules/song"
	"acide/src/utils"
	"net/http"
	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Setup(g *echo.Group) {
	log.Print("Setting up the album module")
	g.Use(utils.Authed)

	g.GET("/", allAlbumsPage)
	g.GET("/:id", albumPage)
}

func allAlbumsPage(c echo.Context) error {
	return utils.RenderTempl(c, http.StatusOK, allAlbumsTempl())
}

func albumPage(c echo.Context) error {
	token, server := utils.Credentials(c)
	albumId := c.Param("id")

	// load album info and song list on the background
	var wg sync.WaitGroup

	var album *utils.Album
	var songs []utils.Song
	var routineErr error = nil

	wg.Add(2)
	go func() {
		defer wg.Done()

		res, err := loadAlbum(token, server, albumId)
		if err != nil {
			routineErr = err
			return
		}

		album = res
	}()
	go func() {
		defer wg.Done()

		res, err := song.LoadSongs(token, server, albumId)
		if err != nil {
			routineErr = err
			return
		}
		songs = res
	}()
	wg.Wait()

	if routineErr != nil {
		return routineErr
	}

	return utils.RenderTempl(c, http.StatusOK, albumTempl(albumId, album, songs))
}
