//Copyright 2019 Chris Wojno
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package param

import (
	"fmt"
	"regexp"
	"strings"
)

// NamedPlaceholderPrefix is the marker that denotes the start of a named placeholder
const NamedPlaceholderPrefix = ":"

type named struct {
	Namer

	// parameters represents the values passed to the object through repeated calls to Set and/or from initialization
	parameters map[string]interface{}

	// vquery is the SQL with the named-placeholders already in the string
	query string

	// cached values after interpolate is run
	// Saves work. Reset both by setting queryNormalized to empty string, this will trigger a full re-run
	// queryNormalized is the `vquery` value transformed back into a driver-specific format
	queryNormalized string

	// orderedNamedParameters is the name of keys to look up in the order that the values of those keys need to appear in the parameterized vquery
	orderedNamedParameters []string
}

// NewNamed creates a new named vquery
// @param vquery is the SQL you wish to execute. When you want to insert a value, use the colon (:) to denote the start of a named parameter. e.g. ":name"
// @return the Queryer-conforming parameterer
// @example
//   vquery: "select * from users where name = :name AND :age = years_old"
//   data: map[string]interface{}{"name": "bob", "age": 21}
//   queries for: users named "bob" who are 21 years old
func NewNamed(query string) Namer {
	return &named{
		query:      query,
		parameters: make(map[string]interface{}),
	}
}

// NewNamedWithData creates a new named vquery and also allows you to pass in named parameters
// This makes one-line vquery-building easier
// @param vquery is the SQL you wish to execute. When you want to insert a value, use the colon (:) to denote the start of a named parameter. e.g. ":name"
// @param data are the key-value pairs for the parameterized values you wish to use with the vquery
// @return the Queryer-conforming parameterer
// @example
//   vquery: "select * from users where name = :name AND :age = years_old"
//   data: map[string]interface{}{"name": "bob", "age": 21}
//   queries for: users named "bob" who are 21 years old
func NewNamedWithData(query string, data map[string]interface{}) Namer {
	return &named{
		query:      query,
		parameters: data,
	}
}

// Set inserts a key-value parameter pair. If it already existed, it is overwritten
func (p *named) Set(key string, value interface{}) {
	p.parameters[key] = value
	p.queryNormalized = ""
}

func (p *named) SQLQuery(strategy InterpolateStrategy) string {
	p.normalizeSQL(strategy)
	return p.queryNormalized
}

func (p *named) Interpolate(strategy InterpolateStrategy) (query string, params []interface{}, err error) {
	p.normalizeSQL(strategy)
	orderedParams := make([]interface{}, 0, len(p.parameters))
	for _, key := range p.orderedNamedParameters {
		if value, ok := p.parameters[key]; !ok {
			// named parameter found, but no mapping for its value was found
			err = &ErrMissingNamedParam{name: key}
			return
		} else {
			orderedParams = append(orderedParams, value)
		}
	}
	return p.queryNormalized, orderedParams, err
}

// normalizeSQL converts the sql-string with the named parameters into the driver-specific format (depending on the strategy)
//
// normalizeSQL stores this value in a cached internal field in case Interpolate is called multiple times
//
// @param strategy is how to insert placeholder for the driver-specific format
func (p *named) normalizeSQL(strategy InterpolateStrategy) {
	if len(p.queryNormalized) == 0 {
		replacementCount := strings.Count(p.query, NamedPlaceholderPrefix)
		p.orderedNamedParameters = make([]string, 0, replacementCount)
		sb := strings.Builder{}
		parts := strings.Split(p.query, NamedPlaceholderPrefix)
		sb.WriteString(parts[0])
		if len(parts) > 1 {
			for i := 1; i < len(parts); i++ {
				match := namedParameterName.FindStringIndex(parts[i])
				paramName := parts[i][match[0]:match[1]]
				p.orderedNamedParameters = append(p.orderedNamedParameters, paramName)
				// remove the name
				sb.WriteString(strategy.InsertPlaceholderIntoSQL())
				sb.WriteString(parts[i][match[1]:])
			}
		}
		p.queryNormalized = sb.String()
	}
}

// ErrMissingNamedParam is a custom error message so we can indicate which key was used that didn't exist
type ErrMissingNamedParam struct {
	name string
}

// Error satisfies the Error interface and prints a lovely message about which key was used but was missing. Very helpful for debugging ;)
func (e ErrMissingNamedParam) Error() string {
	return fmt.Sprintf(`named parameter "%s" was not set to a value`, e.name)
}

// namedParameterName is a pre-compiled regular expression used to find the name of the named parameter AFTER the marker has been identified with colon (:)
var namedParameterName *regexp.Regexp

func init() {
	// looks for the value after the ":". Strings will have already been split, so this value will always be at the beginning of the part. This finds the ":name" where name is the key identifier for the named parameter
	namedParameterName = regexp.MustCompile("^([a-zA-Z_][a-zA-Z0-9_]*)")
}
