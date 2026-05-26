package controllers

import (
	"errors"
	"net/http"
)

var (
	ErrUserNotFound = errors.New(unf)
	jsonErrorMsg    = `{"error":"%s"}`
)

const (
	unf = "user not found"
)

func jsonError(writer http.ResponseWriter, message string, status int) {
	_ = "STUB: not implemented"
	return
}
