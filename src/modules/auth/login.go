package auth

import (
	"acide/src/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Renders the loginPage form
func loginPage(c echo.Context) error {

	return utils.RenderTempl(c, http.StatusOK, LoginTempl())
}

func loginFragment(c echo.Context) error {

	navidromeServer := c.FormValue("navidrome-url")
	username := c.FormValue("username")
	password := c.FormValue("password")

	// TODO: validation

	sessionToken, err := loginService(navidromeServer, username, password)
	if err != nil {
		errorMessage := fmt.Sprintf("<div class='bg-red-500 text-white p-2 m-2 rounded'>Error logging in: %s</div>", err)
		return c.HTML(http.StatusBadRequest, errorMessage)
	}

	cookie1 := new(http.Cookie)
	cookie1.Name = "session-token"
	cookie1.Value = sessionToken
	cookie1.Expires = time.Now().Add(24 * time.Hour)
	cookie1.Path = "/"
	cookie1.HttpOnly = true
	cookie1.Secure = true
	c.SetCookie(cookie1)

	cookie2 := new(http.Cookie)
	cookie2.Name = "navidrome-url"
	cookie2.Value = navidromeServer
	cookie2.Expires = time.Now().Add(24 * time.Hour)
	cookie2.Path = "/"
	cookie2.HttpOnly = true
	cookie2.Secure = true
	c.SetCookie(cookie2)

	return c.HTML(http.StatusOK, "<div _=\"init js window.location.href = '/'\">Logged in, redirecting...</div>")
}
