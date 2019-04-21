package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/jon-whit/go-contacts/interfaces"
	"github.com/jon-whit/go-contacts/models"
)

type ContactsController struct {
	ContactManager interfaces.ContactsManager
}

func (controller *ContactsController) ListUserContacts(w http.ResponseWriter, req *http.Request) {

	userID := chi.URLParam(req, "userid")

	contacts, err := controller.ContactManager.ListUserContacts(userID)
	if err != nil {
		// Handle error
	}

	json.NewEncoder(w).Encode(contacts)
	return
}

func (controller *ContactsController) CreateContact(w http.ResponseWriter, req *http.Request) {

	userID := chi.URLParam(req, "userid")
	contact := models.Contact{}

	// See go-chi/render package
	if err := render.Bind(req, &contact); err != nil {
		// Handle error
	}

	createdContact, err := controller.ContactManager.CreateContact(userID, contact)
	if err != nil {
		// Handle error
	}

	json.NewEncoder(w).Encode(createdContact)
}
