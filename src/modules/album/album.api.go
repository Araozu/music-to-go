package album

import (
	"acide/src/modules/song"
	"acide/src/utils"
	"sync"

	"github.com/labstack/echo/v4"
)

type GetAlbumApiDTO struct {
	Album *utils.Album `json:"album"`
	Songs []utils.Song `json:"songs"`
}

func getAlbumApi(c echo.Context) error {
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
		return c.String(500, routineErr.Error())
	}

	return c.JSON(200, GetAlbumApiDTO{Album: album, Songs: songs})
}
