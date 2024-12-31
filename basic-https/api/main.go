package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

const webPort = 8080

func NewConfig() *Config {
	return &Config{}
}

func main() {
	log.Print("Starting API Server")
	config := NewConfig()

	// Use the routes() method to initialize and return the router
	handler := config.routes()

	// Start the HTTP server with the returned handler
	log.Printf("Starting broker on port:%d\n", webPort)
	port := fmt.Sprintf(":%d", webPort)
	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
