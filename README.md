# go-contacts
go-contacts is a REST API to manage Contacts written in Golang.

## Overview
The Contacts API specification is defined as follows:

| Method | Path                                 | Description                    |
|--------|--------------------------------------|--------------------------------|
| GET    | `/users/:userid/contacts`            | List a user's contacts         |
| POST   | `/users/:userid/contacts`            | Create a contact               |

## Install
Clone the source
    git clone https://github.com/jon-whit/go-contacts

Setup dependencies

    go get -u github.com/go-chi/chi
    go get -u github.com/mattn/go-sqlite3

Setup sqlite data structure

    sqlite3 /var/tmp/go-contacts.db < setup.sql

Test first for your liking

    go test ./... -v

Run the app

    go build && ./go-contacts