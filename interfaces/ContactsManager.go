package interfaces

import "github.com/jon-whit/go-contacts/models"

type ContactsManager interface {
	ListUserContacts(userID string) ([]models.Contact, error)
	CreateContact(userID string, contact models.Contact) (models.Contact, error)
}
