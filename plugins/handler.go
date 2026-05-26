package plugins

import (
	"net/http"
	"plugin"
)

var (
	jsonErrFormat = `{"error": "%s"}`
)

// LoadedPlugin structure for controlling the loaded plugin
type LoadedPlugin struct {
	Loaded bool
	Plugin *plugin.Plugin
}

// PluginFuncReturn structure for holding return value and status of plugin function.
type PluginFuncReturn struct {
	ReturnJson string
	StatusCode int
}

// loadedFunc global variable to control plugins loaded, blocking duplicate loading
var loadedFunc = map[string]LoadedPlugin{}

// loadFunc private func to load and exec OS Library
func loadFunc(fileName, funcName string, r *http.Request) (ret PluginFuncReturn, err error) {
	_ = "STUB: not implemented"
	return *new(PluginFuncReturn), nil
}

// plugin will be loaded only on the first call to the endpoint

// HTTPVars populate

// URL Query populate

// function name: HttpMethod+FunctionName+"Handler" (string sufix)
// standardizing the name of the method that will be invoked we use
// the name Handler as a suffix to identify what will be called in the http

// Exec (call) function name, return string (In case which return status code does not matter)

// It is probable that plugin function return not only json but also status code.

// HandlerPlugin responsible for processing the `.so` function via http protocol
func HandlerPlugin(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Cache arrow if enabled

//nolint
