package album

import (
	"acide/src/utils"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func loadAlbums(token, server string, start, end int) ([]utils.Album, error) {
	var albums []utils.Album
	var error utils.NavError

	client := resty.New()
	response, err := client.R().
		SetHeader("x-nd-authorization", fmt.Sprintf("Bearer %s", token)).
		SetResult(&albums).
		SetError(&error).
		Get(fmt.Sprintf("%s/api/album?_start=%d&_end=%d&_sort=name&_order=ASC", server, start, end))

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, errors.New(error.Error)
	}

	return albums, nil
}

func loadAlbum(token, server, albumId string) (*utils.Album, error) {
	var album utils.Album
	var error utils.NavError

	client := resty.New()
	response, err := client.R().
		SetHeader("x-nd-authorization", fmt.Sprintf("Bearer %s", token)).
		SetResult(&album).
		SetError(&error).
		Get(fmt.Sprintf("%s/api/album/%s", server, albumId))

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, errors.New(error.Error)
	}

	return &album, nil
}
