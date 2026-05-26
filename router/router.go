package router

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni/v3"
)

// GetRouter reagister all routes
// v2: this is not used anywhere, so we can make it private
func GetRouter() *mux.Router { _ = "STUB: not implemented"; return nil }

// can be db specific in the future, there's bellow a proposal
// maybe disable on multiple databases

// multiple DB suggestion:
// router.HandleFunc("/db/{database}/auth", controllers.Auth).Methods("POST")

// breaking change

// router.HandleFunc("/_QUERIES/{database}/{queriesLocation}/{script}", controllers.ExecuteFromScripts)
// if it is windows it should not register the plugin endpoint
// we use go plugin system that does not support windows
// https://github.com/golang/go/issues/19282

// plugins middleware

// Routes for pREST
func Routes() *negroni.Negroni { _ = "STUB: not implemented"; return nil }
