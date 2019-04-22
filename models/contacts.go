package models

import "net/http"

type Contact struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email,omitempty"`
}

// Bind on Contact will run after the unmarshalling is complete, its
// a good time to focus some post-processing after a decoding.
func (c *Contact) Bind(r *http.Request) error {
	return nil
}
