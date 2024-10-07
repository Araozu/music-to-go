package src

import (
	"acide/src/database"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"gorm.io/gorm"
)

type Server struct {
	port int
	db   *gorm.DB
}

// Creates a new Server
func NewServer() *http.Server {
	portEnv := os.Getenv("PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		panic(fmt.Sprintf("Error converting PORT env variable to int.\n%s\n%s", portEnv, err))
	}

	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	// TODO: Register DB schemas here
	// login.SetupSchema(NewServer.db)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
