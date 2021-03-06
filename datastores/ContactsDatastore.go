package datastores

import (
	"github.com/jon-whit/go-contacts/interfaces"
	"github.com/jon-whit/go-contacts/models"
)

type ContactsDatastore struct {
	interfaces.DBHandler
}

func (ds *ContactsDatastore) ListUserContacts(userID string) ([]models.Contact, error) {
	rows, err := ds.Query("SELECT (id, first_name, last_name, phone, email) FROM contacts WHERE user_id=?", userID)
	if err != nil {
		return nil, err
	}

	var contacts []models.Contact
	for rows.Next() {

		var id, firstName, lastName, phone, email string

		err = rows.Scan(&id, &firstName, &lastName, &phone, &email)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, models.Contact{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Phone:     phone,
			Email:     email,
		})
	}

	return contacts, nil
}

func (ds *ContactsDatastore) CreateUserContact(userID string, contact models.Contact) (models.Contact, error) {
	_, err := ds.Execute(`INSERT INTO contacts (user_id, id, first_name, last_name, phone, email) VALUES (?,?,?,?,?,?)`,
		userID, contact.ID, contact.FirstName, contact.LastName, contact.Phone, contact.Email)

	if err != nil {
		return models.Contact{}, err
	}

	return contact, nil
}
