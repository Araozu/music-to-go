package covers

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func loadCover(token, server, albumId string) ([]byte, error) {

	response, err := resty.New().R().
		SetHeader("x-nd-authorization", fmt.Sprintf("Bearer %s", token)).
		Get(fmt.Sprintf(
			"%s/rest/getCoverArt.view?id=%s&u=%s&s=12e7f3&t=%s&v=1.13.0&c=wmusic&size=300",
			server,
			albumId,
			"fernando",
			"d7bbe92d7da363aa202ae16136887adc",
		))

	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		return nil, errors.New("Error fetching image from server")
	}

	return response.Body(), nil
}
