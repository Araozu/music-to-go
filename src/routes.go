package src

import (
	"acide/src/modules/album"
	"acide/src/modules/artist"
	"acide/src/modules/auth"
	"acide/src/modules/covers"
	"acide/src/modules/index"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var dev = os.Getenv("APP_ENV") == "dev"

// Sets up the Echo server, and registers all routes and sub routes
func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CSRF())

	// Static files
	staticFilesFolder := disableCacheInDevMode(http.StripPrefix("/public",
		http.FileServer(http.Dir("public"))))
	e.GET("/public/*", echo.WrapHandler(staticFilesFolder))

	// NOTE: Register subroutes here
	index.SetupRoutes(e.Group(""))
	auth.SetupRoutes(e.Group("/auth"))
	covers.Setup(e.Group("/covers"))
	album.Setup(e.Group("/album"))
	artist.Setup(e.Group("/artist"))

	return e
}

func disableCacheInDevMode(next http.Handler) http.Handler {
	if !dev {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
