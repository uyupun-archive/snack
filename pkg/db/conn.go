package db

import (
	"database/sql"
	"database/sql/driver"

	"entgo.io/ent/dialect"
	"modernc.org/sqlite"

	"snack/pkg/db/ent"
)

type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		conn.Close()
		return conn, err
	}

	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

func init() {
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}

func Connect() (*ent.Client, error) {
	options := []ent.Option{}
	options = append(options, ent.Debug())
	dsn := "file:database.sqlite3?cache=shared"
	conn, err := ent.Open(dialect.SQLite, dsn, options...)
	if err != nil {
		return nil, err
	}
	conn = conn.Debug()

	return conn, nil
}
