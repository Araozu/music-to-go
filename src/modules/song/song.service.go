package song

import (
	"acide/src/utils"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func LoadSongs(token, server, albumId string) ([]utils.Song, error) {
	var songs []utils.Song
	var error utils.NavError

	client := resty.New()
	response, err := client.R().
		SetHeader("x-nd-authorization", fmt.Sprintf("Bearer %s", token)).
		SetResult(&songs).
		SetError(&error).
		Get(fmt.Sprintf("%s/api/song?_end=0&_order=ASC&_sort=album&_start=0&album_id=%s", server, albumId))

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, errors.New(error.Error)
	}

	return songs, nil
}
