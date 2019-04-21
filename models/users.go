package models

type User struct {
	ID       string    `json:"id"`
	Contacts []Contact `json:"contacts"`
}
