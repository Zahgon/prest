package mock

import (
	"database/sql/driver"
)

// mockConn is the mock of driver.Conn
type mockConn struct{}

func (mc *mockConn) Begin() (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}
func (mc *mockConn) Close() (err error) { _ = "STUB: not implemented"; return nil }
func (mc *mockConn) Prepare(q string) (st driver.Stmt, err error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}
func (mc *mockConn) Commit() (err error)   { _ = "STUB: not implemented"; return nil }
func (mc *mockConn) Rollback() (err error) { _ = "STUB: not implemented"; return nil }
