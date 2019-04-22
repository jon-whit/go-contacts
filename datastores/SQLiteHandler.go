package datastores

import (
	"database/sql"

	"github.com/jon-whit/go-contacts/interfaces"
)

type SQLiteHandler struct {
	Conn *sql.DB
}

func (handler *SQLiteHandler) Execute(query string, args ...interface{}) (sql.Result, error) {
	return handler.Conn.Exec(query, args)
}

func (handler *SQLiteHandler) Query(query string, args ...interface{}) (interfaces.DBRow, error) {
	rows, err := handler.Conn.Query(query, args)

	if err != nil {
		return new(SqliteRow), err
	}
	row := new(SqliteRow)
	row.Rows = rows

	return row, nil
}

type SqliteRow struct {
	Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}
