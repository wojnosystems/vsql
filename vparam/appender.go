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
	"strings"
)

const AppenderPlaceholder = "?"

// appender holds the data used for the Query/Exec calls and allows you to build it as you go
type appender struct {
	query
	parameters []interface{}
}

// NewAppend creates a new appending Parameterer in which you can repeatedly append values to the parameter list as desired
func NewAppend(query string) Appender {
	return &appender{
		query:      *newQueryWithSQL(query),
		parameters: make([]interface{}, 0, 1),
	}
}

// NewAppendData creates a new appending Parameterer in which you can repeatedly append values to the parameter list as desired, but this has no query string
func NewAppendData(data ...interface{}) Appender {
	return &appender{
		query:      *newQuery(),
		parameters: data,
	}
}

// NewAppendWithData creates a new appending Parameterer in which you can repeatedly append values to the parameter list as desired
// this version allows you to optionally set a variadic amount of data to append. This makes one-line vquery-building easier
func NewAppendWithData(query string, data ...interface{}) Appender {
	return &appender{
		query:      *newQueryWithSQL(query),
		parameters: data,
	}
}

// Append adds a value to the end of the parameterized values list
// Values are injected into the vquery in the order you call Append
func (p *appender) Append(value interface{}) {
	p.parameters = append(p.parameters, value)
}

func (p *appender) Interpolate(sqlQuery string, strategy interpolation_strategy.InterpolateStrategy) (interpolatedSQLQuery string, params []interface{}, err error) {
	interpolatedQuery := p.SQLQueryInterpolated(strategy)
	if len(p.parameters) != strings.Count(interpolatedQuery, strategy.InsertPlaceholderIntoSQL()) {
		return "", []interface{}{}, ErrParameterPlaceholderMismatch
	}
	return interpolatedQuery, p.parameters, nil
}

func (p *appender) SQLQueryInterpolated(strategy interpolation_strategy.InterpolateStrategy) string {
	return strings.ReplaceAll(p.query.SQLQueryUnInterpolated(), AppenderPlaceholder, strategy.InsertPlaceholderIntoSQL())
}
