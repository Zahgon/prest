package middlewares

import (
	"github.com/urfave/negroni/v3"

	"github.com/prest/prest/v2/cache"
)

// CacheMiddleware simple caching to avoid equal queries to the database
// todo: receive config.PrestConf.Cache to pass to cache.EndpointRules
// this will help removing global config calls
func CacheMiddleware(cfg *cache.Config) negroni.Handler {
	_ = "STUB: not implemented"
	return *new(negroni.Handler)
}

// team will not be used when downloading information, second result ignored
