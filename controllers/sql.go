package controllers

import (
	"net/http"
)

// ExecuteScriptQuery is a function to execute and return result of script query
func ExecuteScriptQuery(rq *http.Request, queriesPath string, script string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExecuteFromScripts is a controller to peform SQL in scripts created by users
func ExecuteFromScripts(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// set db name on ctx

// Cache arrow if enabled

//nolint

// extractHeaders gets from the given request the headers and populate the provided templateData accordingly.
func extractHeaders(rq *http.Request, templateData map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

// extractQueryParameters gets from the given request the query parameters and populate the provided templateData
// accordingly.
func extractQueryParameters(rq *http.Request, templateData map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}
