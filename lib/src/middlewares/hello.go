// nolint
// all plugins must have their package name as `main`
// each plugin is isolated at compile time
package main

import (
	"github.com/urfave/negroni/v3"
)

// BUILD:
// go build -o ./lib/midllewares/hello.so -buildmode=plugin ./lib/src/middlewares/hello.go
func HelloMiddlewareLoad() negroni.Handler { _ = "STUB: not implemented"; return *new(negroni.Handler) }
