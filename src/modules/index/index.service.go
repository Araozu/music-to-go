package index

import (
	"acide/src/utils"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Gets `amount` random albums from the server
func getRandomAlbums(token, server string, amount int) ([]utils.Album, error) {
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
