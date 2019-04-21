# go-contacts
go-contacts is a REST API to manage Contacts written in Golang.

## Overview
The Contacts API specification is defined as follows:

| Method | Path                                 | Description                    |
|--------|--------------------------------------|--------------------------------|
| GET    | `/users/:userid/contacts`            | List a user's contacts         |
| POST   | `/users/:userid/contacts`            | Create a contact               |