package auth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type LoginInputDTO struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginApi(ctx echo.Context) error {
	// extract data from POST JSON
	var input LoginInputDTO
	err := ctx.Bind(&input)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid JSON")
	}

	// perform login

	sessionKey, err := loginService(input.Url, input.Username, input.Password)
	if err != nil {
		return ctx.String(http.StatusUnauthorized, "Wrong credentials")
	}

	// create cookies
	// res.Header().Add("x-auth-token", "prosor-prosor")

	cookie := new(http.Cookie)
	cookie.Name = "session-token"
	cookie.Value = sessionKey
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)

	urlcookie := new(http.Cookie)
	urlcookie.Name = "navidrome-url"
	urlcookie.Value = input.Url
	urlcookie.Path = "/"
	urlcookie.Expires = time.Now().Add(24 * time.Hour)

	ctx.SetCookie(cookie)
	ctx.SetCookie(urlcookie)
	return nil
}
