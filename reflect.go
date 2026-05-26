package arg

import (
	"encoding"
	"reflect"
)

var textUnmarshalerType = reflect.TypeOf([]encoding.TextUnmarshaler{}).Elem()

// cardinality tracks how many tokens are expected for a given spec
//   - zero is a boolean, which does to expect any value
//   - one is an ordinary option that will be parsed from a single token
//   - multiple is a slice or map that can accept zero or more tokens
type cardinality int

const (
	zero cardinality = iota
	one
	multiple
	unsupported
)

func (k cardinality) String() string { _ = "STUB: not implemented"; return "" }

// cardinalityOf returns true if the type can be parsed from a string
func cardinalityOf(t reflect.Type) (cardinality, error) {
	_ = "STUB: not implemented"
	return *new(cardinality), nil
}

// look inside pointer types

// look inside slice and map types

// isBoolean returns true if the type is a boolean or a pointer to a boolean
func isBoolean(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

// isTextUnmarshaler returns true if the type or its pointer implements encoding.TextUnmarshaler
func isTextUnmarshaler(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

// isExported returns true if the struct field name is exported
func isExported(field string) bool { _ = "STUB: not implemented"; return false }

// returns RuneError for empty string or invalid UTF8

// isZero returns true if v contains the zero value for its type
func isZero(v reflect.Value) bool { _ = "STUB: not implemented"; return false }
