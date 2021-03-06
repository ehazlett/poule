package settings

import (
	"fmt"
	"strings"
)

const (
	KeyValuesSeparator = ":"
	ValuesSeparator    = ","
)

// MultiValuedKeys represents a dictionary where each key can have multiple
// values. This is represented:
//
//  - In the CLI using a StringSliceFlag such as:
//	      `--param key=val1,val2,val3`
//
//  - In the configuration file using a dictionary of lists, such as
//        `param: { key: [ val1, val2, val3 ] }`
//
type MultiValuedKeys map[string][]string

// NewMultiValuedKeys returns an empty MultiValuedKeys.
func NewMultiValuedKeys() MultiValuedKeys {
	return MultiValuedKeys{}
}

// NewMultiValuedKeysFromSlice reads a MultiValuedKeys from a collection of
// formatted strings, as typically provided through the command line.
func NewMultiValuedKeysFromSlice(collection []string) (MultiValuedKeys, error) {
	m := NewMultiValuedKeys()
	for _, item := range collection {
		var s []string
		if s = strings.SplitN(item, KeyValuesSeparator, 2); len(s) != 2 {
			return nil, fmt.Errorf("invalid item format %q (expected `key:values`)", item)
		}
		m[s[0]] = strings.Split(s[1], ValuesSeparator)
	}
	return m, nil
}

// ForEach calls the provided `fn` function for each key-value pair.
func (m MultiValuedKeys) ForEach(fn func(key, value string) error) error {
	for key, values := range m {
		for _, value := range values {
			if err := fn(key, value); err != nil {
				return err
			}
		}
	}
	return nil
}

// ToSerializedFormat transforms a MultiValuedKeys into its serialized
// representation, where single values are represented by a string rather than
// a slice of length 1.
func (m MultiValuedKeys) ToSerializedFormat() map[string]interface{} {
	res := map[string]interface{}{}
	for key, values := range m {
		if len(values) == 1 {
			res[key] = values[0]
		} else {
			res[key] = values
		}
	}
	return res
}
