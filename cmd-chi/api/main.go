package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

const webPort = 8080

func main() {
	log.Printf("Starting API server")

	app := &Config{}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: app.routes(),
	}

	log.Printf("Starting broker on port: %d\n", webPort)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
