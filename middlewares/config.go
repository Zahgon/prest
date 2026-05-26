package middlewares

import (
	"github.com/urfave/negroni/v3"
)

var (
	app *negroni.Negroni

	// MiddlewareStack on pREST
	MiddlewareStack []negroni.Handler

	// BaseStack Middlewares
	BaseStack = []negroni.Handler{
		negroni.Handler(negroni.NewRecovery()),
		negroni.Handler(negroni.NewLogger()),
		HandlerSet(),
		SetTimeoutToContext(),
	}
)

func initApp() { _ = "STUB: not implemented"; return }

// GetApp get negroni
func GetApp() *negroni.Negroni { _ = "STUB: not implemented"; return nil }
