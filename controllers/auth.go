package controllers

import (
	"net/http"

	"github.com/prest/prest/v2/controllers/auth"
)

// Response representation
type Response struct {
	LoggedUser interface{} `json:"user_info"`
	Token      string      `json:"token"`
}

// RavensRequest representation
type RavensRequest struct {
	Type       string   `json:"type_of"`
	Subject    string   `json:"subject"`
	Recipients []string `json:"recipients"`
	Sender     string   `json:"sender"`
	SenderName string   `json:"sender_name"`
	Content    string   `json:"content"`
}

// Login representation of data received in authentication
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Token for user
func Token(u auth.User) (t string, err error) {
	_ = "STUB: not implemented"
	// add start time (NotBefore)
	return "", nil
}

// add expiry time in configuration (in minute format, so we support the maximum need)

// TODO: JWT any Algorithm support

// Auth controller
func Auth(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// TODO: form support

// to use body field authentication

//nolint

// to use http basic authentication

// basicPasswordCheck
func basicPasswordCheck(user, password string) (obj auth.User, err error) {
	_ = "STUB: not implemented"
	/**
	table name, fields (user and password) and encryption must be defined in
	the configuration file (toml)
	by default this endpoint will not be available, it is necessary to activate
	in the configuration file
	*/return *new(auth.User), nil
}

// getSelectQuery create the query to authenticate the user
func getSelectQuery() (query string) { _ = "STUB: not implemented"; return "" }

// encrypt will apply the encryption algorithm to the password
func encrypt(password string) (encrypted string) { _ = "STUB: not implemented"; return "" }
