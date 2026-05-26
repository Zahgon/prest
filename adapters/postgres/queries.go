package postgres

import (
	"context"

	"github.com/prest/prest/v2/adapters"
)

// GetScript get SQL template file
func (adapter *Postgres) GetScript(verb, folder, scriptName string) (script string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseScript use values sent by users and add on script
func (adapter *Postgres) ParseScript(scriptPath string, templateData map[string]interface{}) (sqlQuery string, values []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// WriteSQL perform INSERT's, UPDATE's, DELETE's operations
func WriteSQL(sql string, values []interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// WriteSQLCtx perform INSERT's, UPDATE's, DELETE's operations
func WriteSQLCtx(ctx context.Context, sql string, values []interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// ExecuteScripts run sql templates created by users
func (adapter *Postgres) ExecuteScripts(method, sql string, values []interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// ExecuteScriptsCtx run sql templates created by users
func (adapter *Postgres) ExecuteScriptsCtx(ctx context.Context, method, sql string, values []interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}
