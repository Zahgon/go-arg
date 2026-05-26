package arg

import (
	"reflect"
)

// setSliceOrMap parses a sequence of strings into a slice or map. If clear is
// true then any values already in the slice or map are first removed.
func setSliceOrMap(dest reflect.Value, values []string, clear bool) error {
	_ = "STUB: not implemented"
	return nil
}

// setSlice parses a sequence of strings and inserts them into a slice. If clear
// is true then any values already in the slice are removed.
func setSlice(dest reflect.Value, values []string, clear bool) error {
	_ = "STUB: not implemented"
	return nil
}

// clear the slice in case default values exist

// parse the values one-by-one

// setMap parses a sequence of name=value strings and inserts them into a map.
// If clear is true then any values already in the map are removed.
func setMap(dest reflect.Value, values []string, clear bool) error {
	_ = "STUB: not implemented"
	// determine the key and value type
	return nil
}

// clear the slice in case default values exist

// allocate the map if it is not allocated

// parse the values one-by-one

// split at the first equals sign

// parse the key

// parse the value

// add it to the map
