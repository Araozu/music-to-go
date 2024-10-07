package auth

import (
	"log"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Represents a User in the Database
type User struct {
	gorm.Model
	Name     string
	Password string
}

// A package-level database handle.
// It is set up by the [SetupSchema] func,
// which should be called before the webserver
// begins to accept connections
var db *gorm.DB

// Creates all the database tables that this module requires,
// and sets up the database handle for this module
func SetupSchema(dbHandle *gorm.DB) {
	dbHandle.AutoMigrate(&User{})
	db = dbHandle
}

// Registers all the routes that this module (auth) provides
func SetupRoutes(g *echo.Group) {
	log.Print("Setting up the auth module")

	// To just render an HTML template with HTTP 200 status:
	// g.GET("/login", echo.WrapHandler(templ.Handler(LoginTempl())))

	// To include custom rendering logic:
	g.GET("/login", login)
}
