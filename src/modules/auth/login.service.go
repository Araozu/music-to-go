package auth

import (
	"acide/src/utils"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type AuthError struct {
	Error string `json:"error"`
}

// Attempts to login to a navidrome server with the provided credentials.
// Returns the session key if succesful, an error otherwise
func loginService(server, username, password string) (string, error) {
	client := resty.New()

	var loginData utils.AuthSuccess
	var loginError AuthError

	response, err := client.R().
		SetHeader("Content-Type", "").
		SetBody(fmt.Sprintf(`{"username":"%s","password":"%s"}`, username, password)).
		SetResult(&loginData).
		SetError(&loginError).
		Post(fmt.Sprintf("%s/auth/login", server))

	if err != nil {
		return "", err
	}

	if !response.IsSuccess() {
		return "", errors.New(loginError.Error)
	}

	var sessionToken = loginData.Token

	return sessionToken, nil
}
