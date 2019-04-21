package services

import (
	"github.com/jon-whit/go-contacts/interfaces"
	"github.com/jon-whit/go-contacts/models"
)

type ContactsService struct {
	DataAccessor interfaces.ContactsDataAccessor
}

func (service *ContactsService) ListUserContacts(userID string) ([]models.Contact, error) {
	return service.DataAccessor.ListUserContacts(userID)
}

func (service *ContactsService) CreateContact(userID string, contact models.Contact) (models.Contact, error) {
	return service.DataAccessor.CreateUserContact(userID, contact)
}
