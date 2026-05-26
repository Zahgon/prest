package dbtime

import (
	"time"
)

// Time replace MarshalJSON and UnmarshalJSON functions to allow
// compatibility with the Postgresql date format.
type Time struct {
	time.Time
}

const layout = "2006-01-02T15:04:05.999999"

// UnmarshalJSON compatibility with the Postgresql date format
func (t *Time) UnmarshalJSON(b []byte) (err error) { _ = "STUB: not implemented"; return nil }

// MarshalJSON compatibility with the Postgresql date format
func (t Time) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
