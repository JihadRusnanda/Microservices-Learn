package main

import (
	"encoding/json"
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

func (u *Config) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	data := NewUser("jihad1", "stockbil")
	binaryData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = w.Write(binaryData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
