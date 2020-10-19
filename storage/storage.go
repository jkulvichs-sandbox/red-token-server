package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// Storage - Controller for storage
type Storage struct {
	// Type - Storage type: sqlite, etc ...
	Type string
	// DB - Storage's DB handler
	DB *sql.DB
}

// Connect - Binds Storage with any configured DB
func Connect(conf ConfStorage) (*Storage, error) {
	var store *Storage

	switch {
	case conf.SQLite != nil:
		var err error
		store, err = ConnectSQLite(*conf.SQLite)
		if err != nil {
			return nil, fmt.Errorf("can't connect to sqlite storage: %s", err)
		}
	default:
		return nil, fmt.Errorf("at least one of databases must be configured")
	}

	return store, nil
}

// ConnectSQLite - Binds Storage with SQLite DB
func ConnectSQLite(conf ConfSQLiteDataBase) (*Storage, error) {
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=memory", conf.Path)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("can't connect to SQLite: %s", err)
	}

	return &Storage{
		Type: "sqlite",
		DB:   db,
	}, nil
}

// Disconnect - Disconnects from Store
func (storage *Storage) Disconnect() error {
	if err := storage.DB.Close(); err != nil {
		return err
	}
	return nil
}
