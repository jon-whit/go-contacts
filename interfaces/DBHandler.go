package interfaces

import "database/sql"

type DBHandler interface {
	Query(query string, args ...interface{}) (DBRow, error)
	Execute(query string, args ...interface{}) (sql.Result, error)
}

type DBRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}
