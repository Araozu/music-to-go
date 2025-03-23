package auth

import (
	"acide/src/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Renders the loginPage form
func loginPage(c echo.Context) error {
	// if the request has the required cookies, redirect to /
	_, err := c.Cookie("session-token")
	_, err2 := c.Cookie("navidrome-url")
	if err == nil && err2 == nil {
		return c.Redirect(http.StatusFound, "/")
	}

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
	cookie1.SameSite = http.SameSiteStrictMode
	c.SetCookie(cookie1)

	cookie2 := new(http.Cookie)
	cookie2.Name = "navidrome-url"
	cookie2.Value = navidromeServer
	cookie2.Expires = time.Now().Add(24 * time.Hour)
	cookie2.Path = "/"
	cookie2.HttpOnly = true
	cookie2.Secure = true
	cookie2.SameSite = http.SameSiteStrictMode
	c.SetCookie(cookie2)

	return c.HTML(http.StatusOK, "<div _=\"init go to url '/'\">Logged in, redirecting...</div>")
}
