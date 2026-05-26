package plugins

import (
	"github.com/urfave/negroni/v3"
)

// loadedFunc global variable to control plugins loaded, blocking duplicate loading
var loadedMiddlewareFunc = map[string]LoadedPlugin{}

// loadFunc private func to load and exec OS Library
func loadMiddlewareFunc(fileName, funcName string) (handlerFunc negroni.HandlerFunc, err error) {
	_ = "STUB: not implemented"
	// path to plugin file ex: `./libs/middlewares/hello.so`
	return *new(negroni.HandlerFunc), nil
}

// plugin will be loaded only on the first call to the endpoint

// function name: FunctionName+"MiddlewareLoad" (string sufix)
// standardizing the name of the method that will be invoked we use
// the name Handler as a suffix to identify what will be called in the http

// Exec (call) function name, return `negroni.HandlerFunc`

// MiddlewarePlugin responsible for processing the `.so` middleware pattern
/**
example .toml config:
[[pluginmiddlewarelist]]
file = "hello_midlleware.so"
func = "Hello"
*/
func MiddlewarePlugin() negroni.Handler { _ = "STUB: not implemented"; return *new(negroni.Handler) }

// list of plugins configured to be loaded

// negroni not support nil, return empty middleware to continue request
