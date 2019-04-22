package main

import (
	"database/sql"
	"sync"

	"github.com/go-chi/chi"
	"github.com/jon-whit/go-contacts/controllers"
	"github.com/jon-whit/go-contacts/datastores"
	"github.com/jon-whit/go-contacts/services"
)

type ChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	// Create the SQLite DB Handler
	sqlConn, err := sql.Open("sqlite3", "/var/tmp/go-contacts.db")
	if err != nil {
		// handle error
	}
	sqliteHandler := &datastores.SQLiteHandler{
		Conn: sqlConn,
	}

	// Inject all implementations of the interfaces.
	controller := controllers.ContactsController{
		&services.ContactsService{
			DataAccessor: &datastores.ContactsDatastore{
				sqliteHandler,
			},
		},
	}

	// Define and bind the API routes for the Contacts API
	r := chi.NewRouter()
	r.Get("/users/{userid}/contacts", controller.ListUserContacts)
	r.Post("/users/{userid}/contacts", controller.CreateContact)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

// NewChiRouter defines a Singleton, ensuring only a single ChiRouter is created
func NewChiRouter() ChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
