//go:generate mockgen -source ../interfaces/ContactsDataAccessor.go -destination ../mocks/mock_ContactsDataAccessor.go -package mocks ContactsDataAccessor

package interfaces

import "github.com/jon-whit/go-contacts/models"

type ContactsDataAccessor interface {
	ListUserContacts(userID string) ([]models.Contact, error)
	CreateUserContact(userID string, contact models.Contact) (models.Contact, error)
}
