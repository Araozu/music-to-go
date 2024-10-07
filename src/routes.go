package src

import (
	"acide/src/modules/auth"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var dev = os.Getenv("DEV") != ""

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

	// TODO: Register subroutes here
	// login.SetupRoutes(e.Group("/auth"))
	auth.SetupRoutes(e.Group("/auth"))

	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello")
	})

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
