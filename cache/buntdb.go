package cache

import (
	"net/http"

	"github.com/tidwall/buntdb"
)

// BuntConnect connects to database BuntDB - used for caching
func (c *Config) BuntConnect(key string) (db *buntdb.DB, err error) {
	_ = "STUB: not implemented"

	// each url will have its own cache,
	// this will avoid slowing down the cache base
	// it is saved in a file on the file system
	return nil, nil
}

// in case of an error to open buntdb the prestd cache is forced to false

// BuntGet downloads the data - if any - that is in the buntdb (embedded cache database)
// using response.URL.String() as key
func (c Config) BuntGet(key string, w http.ResponseWriter) (cacheExist bool) {
	_ = "STUB: not implemented"
	return false
}

//nolint:errcheck

// BuntSet sets data as cache in buntdb (embedded cache database)
// using response.URL.String() as key
func (c Config) BuntSet(key, value string) { _ = "STUB: not implemented"; return }

//nolint:errcheck

//nolint:errcheck
