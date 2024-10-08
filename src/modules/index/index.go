package index

import (
	"acide/src/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Registers all the routes that this module (auth) provides
func SetupRoutes(g *echo.Group) {
	log.Print("Setting up the index module")

	// To include custom rendering logic:
	g.GET("/", index)
	g.POST("/f/login", loginFragment)
}

func index(c echo.Context) error {
	return utils.RenderTempl(c, http.StatusOK, IndexTempl())
}

func loginFragment(c echo.Context) error {

	navidromeServer := c.FormValue("navidrome-url")
	username := c.FormValue("username")
	password := c.FormValue("password")

	// TODO: validation

	sessionToken, err := login(navidromeServer, username, password)
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

	return c.HTML(http.StatusOK, "wrote some cookies :D")
}
