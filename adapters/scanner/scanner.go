package scanner

import (
	"bytes"
	"errors"
	"reflect"
)

var (
	errPtr      = errors.New("item to input data is not a pointer")
	errUnsupTyp = errors.New("item to input data has an unsupported type")
	errLength   = errors.New("rows returned is not 1")
	supType     = map[reflect.Kind]bool{
		reflect.Slice:  true,
		reflect.Struct: true,
		reflect.Map:    true,
	}
)

func validateType(i interface{}) (ref reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// PrestScanner is a default implementation of adapter.Scanner
type PrestScanner struct {
	Buff    *bytes.Buffer
	Error   error
	IsQuery bool
}

// Scan put prest response into a struct or map
func (p *PrestScanner) Scan(i interface{}) (l int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *PrestScanner) scanQuery(ref reflect.Value, i interface{}) (l int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *PrestScanner) scanNotQuery(ref reflect.Value, i interface{}) (l int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Bytes return prest response in bytes
func (p *PrestScanner) Bytes() (byt []byte) { _ = "STUB: not implemented"; return nil }

// Err return prest response error
func (p *PrestScanner) Err() (err error) { _ = "STUB: not implemented"; return nil }
