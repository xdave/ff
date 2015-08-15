// Copyright 2015 Dave Gradwell
// Under BSD-style license (see LICENSE file)

package ff

import (
	"fmt"
)

// Contains the key and/or value of a command-line parameter
type Param struct {
	Key   interface{}
	Value interface{}
}

// If `key` is not a string, or is nil, it will be ignored
// If `value` is nil, it will be ignored
// Always returns a new Param struct, even if the fields are nil
func NewParam(key, value interface{}) (param Param) {
	switch key.(type) {
	case string:
		param.Key = key
	default:
		param.Key = nil
	}
	param.Value = value
	return
}

// Converts this parameter to a []string{}
// As above, if the `key` is not a string, it won't be added to the slice
// And if `value` is nil, it will not be added to the slice
// Always returns a []string slice, even if empty
func (p Param) Slice() (pair []string) {
	pair = []string{}
	if p.Key != nil {
		pair = append(pair, fmt.Sprintf("-%v", p.Key))
	}
	if p.Value != nil {
		pair = append(pair, fmt.Sprintf("%v", p.Value))
	}
	return
}
