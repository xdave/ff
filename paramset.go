// Copyright 2015 Dave Gradwell
// Under BSD-style license (see LICENSE file)

package ff

// Represents a set of Params
type ParamSet []string

// If no params are passed in, nothing is added.
// If the Slice() representation of a passed-in param is 0, they won't be added
// Always returns a pointer to a new ParamSet, never nil
func NewParamSet(params ...Param) *ParamSet {
	items := &ParamSet{}
	items.Add(params...)
	return items
}

// Just runs len() on the slice
func (p *ParamSet) Len() int {
	return len([]string(*p))
}

// Contains exactly the same functionality as NewParamSet()
// Does not return any value
func (p *ParamSet) Add(params ...Param) {
	for _, param := range params {
		slice := param.Slice()
		if len(slice) > 0 {
			*p = append(*p, slice...)
		}
	}
}

// Always returns a []string slice, even if empty
func (p *ParamSet) Slice() []string {
	return []string(*p)
}
