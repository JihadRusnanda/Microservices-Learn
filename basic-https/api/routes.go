package main

import (
	"net/http"
)

func (u *Config) routes() http.Handler {
	mux := http.NewServeMux()
	// Register route for creating a user
	mux.HandleFunc("/auth-service/create-user", u.CreateUser)
	return mux
}
