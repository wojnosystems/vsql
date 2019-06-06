//Copyright 2019 Chris Wojno
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
// WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS
// OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package vparam

import (
	"github.com/wojnosystems/vsql/interpolation_strategy"
	"regexp"
	"strings"
)

// NamedPlaceholderPrefix is the marker that denotes the start of a named placeholder
const NamedPlaceholderPrefix = ":"

type named struct {
	query

	// parameters represents the values passed to the object through repeated calls to Set and/or from initialization
	parameters map[string]interface{}
}

// NewNamed creates a new named query
// @vparam query is the SQL you wish to execute. When you want to insert a value, use the colon (:) to denote the start of a named parameter. e.g. ":name"
// @return the Queryer-conforming parameterer
// @example
//   query: "select * from users where name = :name AND :age = years_old"
//   data: map[string]interface{}{"name": "bob", "age": 21}
//   queries for: users named "bob" who are 21 years old
func NewNamed(query string) Namer {
	return &named{
		query:      *newQueryWithSQL(query),
		parameters: make(map[string]interface{}),
	}
}

// NewNamedWithData creates a new named query and also allows you to pass in named parameters
// This makes one-line query-building easier
// @vparam query is the SQL you wish to execute. When you want to insert a value, use the colon (:) to denote the start of a named parameter. e.g. ":name"
// @vparam data are the key-value pairs for the parameterized values you wish to use with the query
// @return the Queryer-conforming parameterer
// @example
//   query: "select * from users where name = :name AND :age = years_old"
//   data: map[string]interface{}{"name": "bob", "age": 21}
//   queries for: users named "bob" who are 21 years old
func NewNamedWithData(query string, data map[string]interface{}) Namer {
	return &named{
		query:      *newQueryWithSQL(query),
		parameters: data,
	}
}

// NewNamedData creates a new named query and also allows you to pass in named parameters, but not a string query
// This makes one-line query-building easier
// @vparam data are the key-value pairs for the parameterized values you wish to use with the query
// @return the Queryer-conforming parameterer
// @example
//   query: "select * from users where name = :name AND :age = years_old"
//   data: map[string]interface{}{"name": "bob", "age": 21}
//   queries for: users named "bob" who are 21 years old
func NewNamedData(data map[string]interface{}) Namer {
	return &named{
		query:      *newQuery(),
		parameters: data,
	}
}

// Set inserts a key-value parameter pair. If it already existed, it is overwritten
func (p *named) Set(key string, value interface{}) {
	p.parameters[key] = value
}

func (p *named) Interpolate(sqlQuery string, strategy interpolation_strategy.InterpolateStrategy) (interpolatedSQLQuery string, params []interface{}, err error) {
	var orderedNamedParameters []string
	interpolatedSQLQuery, orderedNamedParameters = normalizeSQL(sqlQuery, strategy)
	orderedParams := make([]interface{}, 0, len(p.parameters))
	for _, key := range orderedNamedParameters {
		if value, ok := p.parameters[key]; !ok {
			// named parameter found, but no mapping for its value was found
			err = &ErrMissingNamedParam{name: key}
			return
		} else {
			orderedParams = append(orderedParams, value)
		}
	}
	return interpolatedSQLQuery, orderedParams, err
}

// normalizeSQL converts the sql-string with the named parameters into the driver-specific format (depending on the strategy)
//
// @param unInterpolatedSQLQuery is the query from the Queryer object's SQLQueryUnInterpolated(). This should be the query as provided by the developer with placeholders for the named variables intead of question marks or ordinal parameters
// @param strategy is how to insert placeholder for the driver-specific format
// @return interpolatedSQLQuery is the query with the InterpolateStrategy parameters instead of the names of the parameter placeholders
// @return orderedNamedParameters is the order in which parameters were encountered in the unInterpolatedSQLQuery. These are just tokens and not the actual values
func normalizeSQL(unInterpolatedSQLQuery string, strategy interpolation_strategy.InterpolateStrategy) (interpolatedSQLQuery string, orderedNamedParameters []string) {
	replacementCount := strings.Count(unInterpolatedSQLQuery, NamedPlaceholderPrefix)
	orderedNamedParameters = make([]string, 0, replacementCount)
	sb := strings.Builder{}
	parts := strings.Split(unInterpolatedSQLQuery, NamedPlaceholderPrefix)
	sb.WriteString(parts[0])
	if len(parts) > 1 {
		for i := 1; i < len(parts); i++ {
			match := namedParameterName.FindStringIndex(parts[i])
			paramName := parts[i][match[0]:match[1]]
			orderedNamedParameters = append(orderedNamedParameters, paramName)
			// remove the name
			sb.WriteString(strategy.InsertPlaceholderIntoSQL())
			sb.WriteString(parts[i][match[1]:])
		}
	}
	interpolatedSQLQuery = sb.String()
	return
}

// namedParameterName is a pre-compiled regular expression used to find the name of the named parameter AFTER the marker has been identified with colon (:)
var namedParameterName *regexp.Regexp

func init() {
	// looks for the value after the ":". Strings will have already been split, so this value will always be at the beginning of the part. This finds the ":name" where name is the key identifier for the named parameter
	namedParameterName = regexp.MustCompile("^([a-zA-Z_][a-zA-Z0-9_]*)")
}
