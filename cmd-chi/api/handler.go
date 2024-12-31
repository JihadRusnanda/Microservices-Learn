package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}

func (a *Config) CreateUser(w http.ResponseWriter, r *http.Request) {
	data := NewUser("jihad", "stockbil")
	out, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//w.Write([]byte("data"))

	_, err = w.Write(out)
	if err != nil {
		log.Fatal(err)
	}
}
