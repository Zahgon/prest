package ident

import (
	"regexp"
)

var re = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*(\.[A-Za-z_][A-Za-z0-9_]*)*$`)

// IsValid reports whether s is a valid SQL identifier or dotted identifier path.
func IsValid(s string) bool { _ = "STUB: not implemented"; return false }

// IsSafeSegment reports whether s is a safe, single identifier segment for path params
// like database, schema, or table. It allows letters, digits, underscore and hyphen,
// PostgreSQL identifiers don't support. These must be quoted when used in SQL.
// with length up to 63, and disallows dots and quotes.
func IsSafeSegment(s string) bool { _ = "STUB: not implemented"; return false }

// Quote validates and returns a safely quoted identifier path like "a"."b".
func Quote(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// SplitAndValidateCSV splits a comma-separated list and validates each identifier.
func SplitAndValidateCSV(s string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
