package middlewares

import (
	"net/http"
	"net/http/httptest"
)

var (
	ErrXMLBadRequest = `<?xml version="1.0" encoding="utf-8"?>
<errors xmlns="http://schemas.google.com/g/2005">
  <error>
    <reason>internal</reason>
    <internalReason>%s</internalReason>
  </error>
  <code>400</code> 
</errors>`
)

func getVars(path string) (paths map[string]string) { _ = "STUB: not implemented"; return nil }

func permissionByMethod(method string) (permission string) { _ = "STUB: not implemented"; return "" }

func renderFormat(w http.ResponseWriter, recorder *httptest.ResponseRecorder, format string) {
	_ = "STUB: not implemented"
	return
}

var defaultAllowMethods = []string{
	"GET",
	"POST",
	"PUT",
	"PATCH",
	"DELETE",
	"OPTIONS",
}

const (
	headerAllowOrigin      = "Access-Control-Allow-Origin"
	headerAllowCredentials = "Access-Control-Allow-Credentials"
	headerAllowHeaders     = "Access-Control-Allow-Headers"
	headerAllowMethods     = "Access-Control-Allow-Methods"
	headerOrigin           = "Origin"
)

func checkCors(r *http.Request, origin []string) (allowed bool) {
	_ = "STUB: not implemented"
	return false
}

// MatchURL matches the given url with a whitelist from config.core
func MatchURL(url string) (match bool, err error) { _ = "STUB: not implemented"; return false, nil }
