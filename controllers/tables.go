package controllers

import (
	"net/http"
)

// GetTables list all (or filter) tables
func GetTables(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// GetTablesByDatabaseAndSchema list all (or filter) tables based on database and schema
func GetTablesByDatabaseAndSchema(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// validate safe segments for path params

// set db name on ctx

// send ctx to query the proper DB

// SelectFromTables perform select in database
func SelectFromTables(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// validate path identifiers early using safe segments policy

// Get user info from token

// get selected columns, "*" if empty "_columns"

// sql query formatting if there is a distinct rule

// sql query formatting if there is a count rule

// _count_first: query string

// count returns a list, passing this parameter will return the first
// record as a non-list object

// sql query formatting if there is a join (inner, left, ...) rule

// sql query formatting if there is a where rule

// sql query formatting if there is a groupby rule

// sql query formatting if there is a orderby rule

// sql query formatting if there is a paganate rule

// QueryCount returns the first record of the postgresql return as a non-list object

// Cache arrow if enabled

//nolint

// InsertInTables perform insert in specific table
func InsertInTables(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// validate safe segments for path params

// set db name on ctx

// BatchInsertInTables perform insert in specific table from a batch request
func BatchInsertInTables(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// validate safe segments for path params

// set db name on ctx

// DeleteFromTable perform delete sql
func DeleteFromTable(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// validate safe segments for path params

// validate safe segments for path params

// UpdateTable perform update table
func UpdateTable(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// validate safe segments for path params

// placeholder id

// ShowTable show information from table
func ShowTable(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// validate safe segments for path params

// set db name on ctx
