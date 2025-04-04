package auth

import (
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// A package-level database handle.
// It is set up by the [SetupSchema] func,
// which should be called before the webserver
// begins to accept connections
var db *gorm.DB

// Creates all the database tables that this module requires,
// and sets up the database handle for this module
func SetupSchema(dbHandle *gorm.DB) {
}

// Registers all the routes that this module (auth) provides
func SetupRoutes(g *echo.Group) {
	log.Print("Setting up the auth module")

	// To just render an HTML template with HTTP 200 status:
	// g.GET("/login", echo.WrapHandler(templ.Handler(LoginTempl())))

	// To include custom rendering logic:
	g.GET("/", loginPage)
	g.POST("/f/login", loginFragment)
}

func SetupApiRoutes(g *echo.Group) {
	log.Print("Setting up the auth API module")

	g.POST("/login", LoginApi)
}
