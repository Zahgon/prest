package postgres

import (
	"github.com/prest/prest/v2/adapters/postgres/internal/connection"

	"github.com/jmoiron/sqlx"
)

// GetURI postgres connection URI
func GetURI(DBName string) string { _ = "STUB: not implemented"; return "" }

// Get get postgres connection
func Get() (*sqlx.DB, error) {
	_ = "STUB: not implemented"
	return nil,

		// GetPool of connection
		nil
}

func GetPool() *connection.Pool { _ = "STUB: not implemented"; return nil }

// AddDatabaseToPool add connection to pool
func AddDatabaseToPool(name string) (*sqlx.DB, error) { _ = "STUB: not implemented"; return nil, nil }

// MustGet get postgres connection
func MustGet() *sqlx.DB { _ = "STUB: not implemented"; return nil }

// SetDatabase set current database in use
// todo: remove when ctx is fully implemented
func SetDatabase(name string) { _ = "STUB: not implemented"; return }

// GetDatabase get current database in use
func GetDatabase() string { _ = "STUB: not implemented"; return "" }
