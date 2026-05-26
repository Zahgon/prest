package controllers

import (
	"context"
	"net/http"
)

type CheckList []func(context.Context) error

var DefaultCheckList = CheckList{
	CheckDBHealth,
}

func CheckDBHealth(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func WrappedHealthCheck(checks CheckList) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
