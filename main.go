package main

import (
	"acide/src"
	"fmt"
	"log"
)

func main() {
	server := src.NewServer()
	log.Print("HTTP server started")
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
