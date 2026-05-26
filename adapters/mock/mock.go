package mock

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"net/http"
	"sync"
	"testing"

	"github.com/prest/prest/v2/adapters"
)

// Item mock
type Item struct {
	Body    []byte
	Error   error
	IsCount bool
}

// Mock adapter
type Mock struct {
	mtx   *sync.RWMutex
	t     *testing.T
	conns map[string]*mockConn
	Items []Item
}

var _ adapters.Adapter = (*Mock)(nil) // Verify that Mock implements Adapter.

// New mock
func New(t *testing.T) (m *Mock) { _ = "STUB: not implemented"; return nil }

// Open makes Mock implement driver.Driver
func (m *Mock) Open(dsn string) (c driver.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func (m *Mock) validate() { _ = "STUB: not implemented"; return }

func (m *Mock) perform(query bool) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// TablePermissions mock
func (m *Mock) TablePermissions(table string, op string, userName string) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// If userName is empty, means use table access.

// currently, access is granted to all users based on the table settings.
// if it is later discovered that there are specific permission settings for an individual user,
// then the latter settings should be applied.

// GetScript mock
func (m *Mock) GetScript(verb string, folder string, scriptName string) (script string, err error) {
	_ = "STUB: not implemented"

	// ParseScript mock
	return "", nil
}

func (m *Mock) ParseScript(scriptPath string, data map[string]interface{}) (sqlQuery string, values []interface{}, err error) {
	_ = "STUB: not implemented"

	// ExecuteScripts mock
	return "", nil, nil
}

func (m *Mock) ExecuteScripts(method string, sql string, values []interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"

	// ExecuteScripts mock
	return *new(adapters.Scanner)
}

func (m *Mock) ExecuteScriptsCtx(ctx context.Context, method string, sql string, values []interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"

	// WhereByRequest mock
	return *new(adapters.Scanner)
}

func (m *Mock) WhereByRequest(r *http.Request, initialPlaceholderID int) (whereSyntax string, values []interface{}, err error) {
	_ = "STUB: not implemented"

	// ReturningByRequest mock
	return "", nil, nil
}

func (m *Mock) ReturningByRequest(r *http.Request) (ReturningSyntax string, err error) {
	_ = "STUB: not implemented"

	// DatabaseClause mock
	return "", nil
}

func (m *Mock) DatabaseClause(req *http.Request) (query string, hasCount bool) {
	_ = "STUB: not implemented"
	return "", false
}

// OrderByRequest mock
func (m *Mock) OrderByRequest(r *http.Request) (values string, err error) {
	_ = "STUB: not implemented"

	// PaginateIfPossible mock
	return "", nil
}

func (m *Mock) PaginateIfPossible(r *http.Request) (paginatedQuery string, err error) {
	_ = "STUB: not implemented"

	// GetTransaction mock
	return "", nil
}

func (m *Mock) GetTransaction() (tx *sql.Tx, err error) { _ = "STUB: not implemented"; return nil, nil }

// GetTransactionCtx mock
func (m *Mock) GetTransactionCtx(ctx context.Context) (tx *sql.Tx, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Query mock
func (m *Mock) Query(SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// QueryCtx mock
func (m *Mock) QueryCtx(ctx context.Context, SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// SchemaClause mock
func (m *Mock) SchemaClause(req *http.Request) (query string, hasCount bool) {
	_ = "STUB: not implemented"
	return "", false
}

// FieldsPermissions mock
func (m *Mock) FieldsPermissions(r *http.Request, table string, op string, userName string) (fields []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SelectFields mock
func (m *Mock) SelectFields(fields []string) (sql string, err error) {
	_ = "STUB: not implemented"

	// CountByRequest mock
	return "", nil
}

func (m *Mock) CountByRequest(req *http.Request) (countQuery string, err error) {
	_ = "STUB: not implemented"

	// JoinByRequest mock
	return "", nil
}

func (m *Mock) JoinByRequest(r *http.Request) (values []string, err error) {
	_ = "STUB: not implemented"

	// GroupByClause mock
	return nil, nil
}

func (m *Mock) GroupByClause(r *http.Request) (groupBySQL string) {
	_ = "STUB: not implemented"

	// QueryCount mock
	return ""
}

func (m *Mock) QueryCount(SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// QueryCountCtx mock
func (m *Mock) QueryCountCtx(ctx context.Context, SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// ParseInsertRequest mock
func (m *Mock) ParseInsertRequest(r *http.Request) (colsName string, colsValue string, values []interface{}, err error) {
	_ = "STUB: not implemented"

	// Insert mock
	return "", "", nil, nil
}

func (m *Mock) Insert(SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// Insert mock
func (m *Mock) InsertCtx(ctx context.Context, SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// InsertWithTransaction mock
func (m *Mock) InsertWithTransaction(tx *sql.Tx, SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// Delete mock
func (m *Mock) Delete(SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// DeleteCtx mock
func (m *Mock) DeleteCtx(ctx context.Context, SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// DeleteWithTransaction mock
func (m *Mock) DeleteWithTransaction(tx *sql.Tx, SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// SetByRequest mock
func (m *Mock) SetByRequest(r *http.Request, initialPlaceholderID int) (setSyntax string, values []interface{}, err error) {
	_ = "STUB: not implemented"

	// Update mock
	return "", nil, nil
}

func (m *Mock) Update(SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// UpdateCtx mock
func (m *Mock) UpdateCtx(ctx context.Context, SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// UpdateWithTransaction mock
func (m *Mock) UpdateWithTransaction(tx *sql.Tx, SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// DistinctClause mock
func (m *Mock) DistinctClause(r *http.Request) (distinctQuery string, err error) {
	_ = "STUB: not implemented"

	// SetDatabase mock
	return "", nil
}

func (m *Mock) SetDatabase(name string) {
	_ = "STUB: not implemented"

	// SelectSQL mock
	return
}

func (m *Mock) SelectSQL(selectStr string, database string, schema string, table string) (s string) {
	_ = "STUB: not implemented"

	// InsertSQL mock
	return ""
}

func (m *Mock) InsertSQL(database string, schema string, table string, names string, placeholders string) (s string) {
	_ = "STUB: not implemented"

	// DeleteSQL mock
	return ""
}

func (m *Mock) DeleteSQL(database string, schema string, table string) (s string) {
	_ = "STUB: not implemented"

	// UpdateSQL mock
	return ""
}

func (m *Mock) UpdateSQL(database string, schema string, table string, setSyntax string) (s string) {
	_ = "STUB: not implemented"

	// DatabaseWhere mock
	return ""
}

func (m *Mock) DatabaseWhere(requestWhere string) (whereSyntax string) {
	_ = "STUB: not implemented"

	// DatabaseOrderBy mock
	return ""
}

func (m *Mock) DatabaseOrderBy(order string, hasCount bool) (orderBy string) {
	_ = "STUB: not implemented"

	// SchemaOrderBy mock
	return ""
}

func (m *Mock) SchemaOrderBy(order string, hasCount bool) (orderBy string) {
	_ = "STUB: not implemented"

	// TableClause mock
	return ""
}

func (m *Mock) TableClause() (query string) {
	_ = "STUB: not implemented"

	// TableWhere mock
	return ""
}

func (m *Mock) TableWhere(requestWhere string) (whereSyntax string) {
	_ = "STUB: not implemented"

	// TableOrderBy mock
	return ""
}

func (m *Mock) TableOrderBy(order string) (orderBy string) {
	_ = "STUB: not implemented"

	// SchemaTablesClause mock
	return ""
}

func (m *Mock) SchemaTablesClause() (query string) {
	_ = "STUB: not implemented"

	// SchemaTablesWhere mock
	return ""
}

func (m *Mock) SchemaTablesWhere(requestWhere string) (whereSyntax string) {
	_ = "STUB: not implemented"

	// SchemaTablesOrderBy mock
	return ""
}

func (m *Mock) SchemaTablesOrderBy(order string) (orderBy string) {
	_ = "STUB: not implemented"

	// ParseBatchInsertRequest mocl
	return ""
}

func (m *Mock) ParseBatchInsertRequest(r *http.Request) (colsName string, placeholders string, values []interface{}, err error) {
	_ = "STUB: not implemented"

	// BatchInsertValues mock
	return "", "", nil, nil
}

func (m *Mock) BatchInsertValues(SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// BatchInsertValuesCtx mock
func (m *Mock) BatchInsertValuesCtx(ctx context.Context, SQL string, params ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// BatchInsertCopy mock
func (m *Mock) BatchInsertCopy(dbname, schema, table string, keys []string, values ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// BatchInsertCopyCtx mock
func (m *Mock) BatchInsertCopyCtx(ctx context.Context, dbname, schema, table string, keys []string, values ...interface{}) (sc adapters.Scanner) {
	_ = "STUB: not implemented"
	return *new(adapters.Scanner)
}

// ShowTable shows table structure
func (m *Mock) ShowTable(schema, table string) (sc adapters.Scanner) {
	_ = "STUB: not implemented"

	// ShowTableCtx shows table structure
	return *new(adapters.Scanner)
}

func (m *Mock) ShowTableCtx(ctx context.Context, schema, table string) (sc adapters.Scanner) {
	_ = "STUB: not implemented"

	// AddItem on mock object
	return *new(adapters.Scanner)
}

func (m *Mock) AddItem(body []byte, err error, isCount bool) { _ = "STUB: not implemented"; return }

// GetDatabase ron mock db
func (m *Mock) GetDatabase() (db string) { _ = "STUB: not implemented"; return "" }
