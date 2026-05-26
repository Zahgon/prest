package config

import (
	"errors"

	"github.com/prest/prest/v2/adapters"
	"github.com/prest/prest/v2/cache"

	"log/slog"
)

const (
	jsonAggDefault = "jsonb_agg"
	jsonAgg        = "json_agg"
)

// TablesConf informations

type TablesConf struct {
	Name        string   `mapstructure:"name"`
	Permissions []string `mapstructure:"permissions"`
	Fields      []string `mapstructure:"fields"`
}

type UsersConf struct {
	Name   string `mapstructure:"name"`
	Tables []TablesConf
}

// AccessConf informations
type AccessConf struct {
	Restrict    bool
	IgnoreTable []string
	Tables      []TablesConf
	Users       []UsersConf
}

// ExposeConf (expose data) information
type ExposeConf struct {
	Enabled         bool
	DatabaseListing bool
	SchemaListing   bool
	TableListing    bool
}

type PluginMiddleware struct {
	File string
	Func string
}

// Prest basic config
type Prest struct {
	AuthEnabled          bool
	AuthSchema           string
	AuthTable            string
	AuthUsername         string
	AuthPassword         string
	AuthEncrypt          string
	AuthMetadata         []string
	AuthType             string
	HTTPHost             string // HTTPHost Declare which http address the PREST used
	HTTPPort             int    // HTTPPort Declare which http port the PREST used
	HTTPTimeout          int
	PGHost               string
	PGPort               int
	PGUser               string
	PGPass               string
	PGDatabase           string
	PGURL                string
	PGSSLMode            string
	PGSSLCert            string
	PGSSLKey             string
	PGSSLRootCert        string
	ContextPath          string
	PGMaxIdleConn        int
	PGMaxOpenConn        int
	PGConnTimeout        int
	PGCache              bool
	JWTKey               string
	JWTAlgo              string
	JWTWellKnownURL      string
	JWTJWKS              string
	JWTWhiteList         []string
	JSONAggType          string
	MigrationsPath       string
	QueriesPath          string
	AccessConf           AccessConf
	ExposeConf           ExposeConf
	CORSAllowOrigin      []string
	CORSAllowHeaders     []string
	CORSAllowMethods     []string
	CORSAllowCredentials bool
	Debug                bool
	Adapter              adapters.Adapter
	EnableDefaultJWT     bool
	SingleDB             bool
	HTTPSMode            bool
	HTTPSCert            string
	HTTPSKey             string
	Cache                cache.Config
	PluginPath           string
	PluginMiddlewareList []PluginMiddleware
	Logger               *slog.Logger
}

const defaultCacheDir = "./"

var (
	// PrestConf config variable
	PrestConf      *Prest
	configFile     string
	defaultCfgFile = "./prest.toml"
)

// Load configuration
func Load() { _ = "STUB: not implemented"; return }

// ignore cache if disabled

func viperCfg() { _ = "STUB: not implemented"; return }

// avoids db memory leak on req timeout

// todo: replace this with prefer, will need to replace lib/pq
// https://github.com/jackc/pgx/blob/47d631e34be7128997a0aa89b75885cc4ad4c82e/pgconn/config.go#L218

func getPrestConfFile(prestConf string) string { _ = "STUB: not implemented"; return "" }

// Parse pREST config
// todo: split config onto methods to simplify this
func Parse(cfg *Prest) { _ = "STUB: not implemented"; return }

// table access config

// plugin middleware list config

// parseDatabaseURL tries to get from URL the DB configs
func parseDatabaseURL(cfg *Prest) { _ = "STUB: not implemented"; return }

// Parser PG URL, get database connection via string URL

// ErrJWTDefaultEnabledNoKey is returned when the default JWT middleware is
// enabled but no verification material (HMAC key, JWKS or .well-known URL) was
// provided. This guards against accidentally serving requests with an empty
// HMAC key, which would let any client forge bearer tokens. See GHSA-fj7v-859r-2fm4.
var ErrJWTDefaultEnabledNoKey = errors.New(
	"jwt.default is enabled but no verification material was provided " +
		"(set jwt.key, jwt.jwks or jwt.wellknownurl, or disable jwt.default)")

// ErrAuthEnabledNoJWTKey is returned when basic auth is enabled but jwt.key
// is empty. AuthMiddleware uses the same []byte(JWTKey) to verify HS256
// tokens, so an empty key opens the same auth-bypass as the default JWT
// middleware. See GHSA-fj7v-859r-2fm4.
var ErrAuthEnabledNoJWTKey = errors.New(
	"auth.enabled is true but jwt.key is empty (required to verify HS256 tokens)")

// ValidateJWTConfig fails fast when either of the JWT-validating middlewares
// would be installed without any verification material:
//
//   - The default JWT middleware (jwt.default = true) requires jwt.key, a
//     JWKS, or a .well-known URL.
//   - AuthMiddleware (auth.enabled = true) verifies HS256 tokens with
//     jwt.key, so an empty key is unsafe.
//
// The default JWT path also bypasses when Debug is true, so we mirror that
// rule here to avoid blocking debug-mode startups.
//
// Call this from binary entrypoints before serving requests; tests that
// exercise Load() without setting JWT material rely on the middleware-level
// guards (middlewares.JwtMiddleware, middlewares.AuthMiddleware) to fail
// closed at request time.
func ValidateJWTConfig(cfg *Prest) error { _ = "STUB: not implemented"; return nil }

// fetchJWKS tries to get the JWKS from the URL in the config
func fetchJWKS(cfg *Prest) { _ = "STUB: not implemented"; return }

// Call provider to obtain .well-known config

//Retrieve the JWKS from the endpoint

//Convert set to json string

func portFromEnv(cfg *Prest) { _ = "STUB: not implemented"; return }

// cloud factor support: https://help.heroku.com/PPBPA231/how-do-i-use-the-port-environment-variable-in-container-based-apps

// getJSONAgg identifies which json aggregation function will be used,
// support `jsonb` and `json`; `jsonb` is the default value
//
// https://www.postgresql.org/docs/9.5/functions-aggregate.html
func getJSONAgg() (config string) { _ = "STUB: not implemented"; return "" }

func parseDBConfig(cfg *Prest) { _ = "STUB: not implemented"; return }

// cloud factor support: https://devcenter.heroku.com/changelog-items/438

func loadCacheConfig(cfg *Prest) { _ = "STUB: not implemented"; return }

// cache endpoints config

func parseAuthConfig(cfg *Prest) { _ = "STUB: not implemented"; return }

func parseHTTPConfig(cfg *Prest) { _ = "STUB: not implemented"; return }
