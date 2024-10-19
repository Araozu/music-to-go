package album

import (
	"acide/src/modules/song"
	"acide/src/utils"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type ClientSong struct {
	Title   string `json:"title"`
	Artist  string `json:"artist"`
	AlbumId string `json:"albumId"`
	Album   string `json:"album"`
	SongId  string `json:"songId"`
}

func Setup(g *echo.Group) {
	log.Print("Setting up the album module")
	g.Use(utils.Authed)

	g.GET("/", allAlbumsPage)
	g.GET("/:id", albumPage)
}

func allAlbumsPage(c echo.Context) error {
	// if there's a search query, do that
	searchQuery := c.QueryParam("s")
	isHtmxRequest := c.Request().Header.Get("HX-Request") == "true"
	token, server := utils.Credentials(c)

	// if searchQuery is empty, this will get the first 30 albums
	albums, err := searchAlbums(token, searchQuery, server, 0, 30)

	if err != nil {
		return err
	}

	if isHtmxRequest {
		// return just a fragment
		return utils.RenderTempl(c, http.StatusOK, albumsFragment(albums))
	} else {
		// return a full-blown html page
		return utils.RenderTempl(c, http.StatusOK, allAlbumsTempl(albums))
	}
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

	// convert the song list to json
	clientSons := make([]ClientSong, len(songs))
	for i, song := range songs {
		clientSons[i] = ClientSong{
			Title:   song.Title,
			Artist:  song.Artist,
			AlbumId: album.ID,
			Album:   album.Name,
			SongId:  song.ID,
		}
	}
	clientSongsJson, err := json.Marshal(clientSons)
	if err != nil {
		log.Printf("Error marshaling clientSongs: %s", err)
		return err
	}

	return utils.RenderTempl(c, http.StatusOK, albumTempl(albumId, album, songs, string(clientSongsJson)))
}
