package datastores

import (
	"errors"

	"github.com/jon-whit/go-contacts/models"
)

type ContactsDatastore struct {
	interfaces.DbHandler
}

func (ds *ContactsDatastore) ListUserContacts(userID string) ([]models.Contact, error) {
	return nil, errors.New("Not Implemented Yet")
}

func (ds *ContactsDatastore) CreateUserContact(userID string, contact models.Contact) (models.Contact, error) {
	return models.Contact{}, errors.New("Not Implemented Yet")
}
