package album

import (
	"acide/src/utils"
	"errors"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

func searchAlbums(token, query, server string, start, end int) ([]utils.Album, error) {
	var albums []utils.Album
	var error utils.NavError

	encodedQuery := url.QueryEscape(query)

	client := resty.New()
	response, err := client.R().
		SetHeader("x-nd-authorization", fmt.Sprintf("Bearer %s", token)).
		SetResult(&albums).
		SetError(&error).
		Get(fmt.Sprintf("%s/api/album?_start=%d&_end=%d&_sort=name&_order=ASC&name=%s", server, start, end, encodedQuery))

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

// Gets `amount` random albums from the server
func GetRandomAlbums(token, server string, amount int) ([]utils.Album, error) {
	var albums []utils.Album
	var error utils.NavError

	client := resty.New()

	response, err := client.R().
		SetHeader("x-nd-authorization", fmt.Sprintf("Bearer %s", token)).
		SetResult(&albums).
		SetError(&error).
		Get(fmt.Sprintf("%s/api/album?_end=%d&_order=DESC&_sort=random&_start=0", server, amount))

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, errors.New(fmt.Sprintf("Error getting albums: %s", error.Error))
	}

	return albums, nil
}
