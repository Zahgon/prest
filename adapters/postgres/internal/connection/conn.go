package connection

import (
	"sync"

	"github.com/jmoiron/sqlx"
	// Used pg drive on sqlx
	_ "github.com/lib/pq"
)

var (
	pool         *Pool
	currDatabase string
)

// Pool struct
type Pool struct {
	Mtx *sync.Mutex
	DB  map[string]*sqlx.DB
}

// GetURI postgres connection URI
func GetURI(DBName string) string { _ = "STUB: not implemented"; return "" }

// Get get Postgres connection adding it to the pool if needed
func Get() (*sqlx.DB, error) { _ = "STUB: not implemented"; return nil, nil }

// Connection is already in the pool

// Connection is not in the pool, add it

// GetFromPool tries to get the db name from the db pool
// will return error if not found
func GetFromPool(dbName string) (*sqlx.DB, error) { _ = "STUB: not implemented"; return nil, nil }

// GetPool of connection
func GetPool() *Pool { _ = "STUB: not implemented"; return nil }

func getDatabaseFromPool(name string) *sqlx.DB { _ = "STUB: not implemented"; return nil }

// AddDatabaseToPool create and add connection to the pool
func AddDatabaseToPool(name string) (*sqlx.DB, error) { _ = "STUB: not implemented"; return nil, nil }

// MustGet get postgres connection
func MustGet() *sqlx.DB { _ = "STUB: not implemented"; return nil }

// SetDatabase set current database in use
func SetDatabase(name string) { _ = "STUB: not implemented"; return }

// GetDatabase get current database in use
func GetDatabase() string { _ = "STUB: not implemented"; return "" }
