package main

import (
	"log"
)

func main() {
	// Create the server.
	if err := StartServer(); err == nil {
		log.Println("Starting chat server...")
	} else {
		log.Fatalf("Cannot initialize banking server: %s", err)
	}
}
