//Copyright 2019 Chris Wojno
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package aggregator

import (
	"context"
	"database/sql"
	"github.com/wojnosystems/vsql/param"
	"github.com/wojnosystems/vsql/ulong"
	"github.com/wojnosystems/vsql/vquery"
	"github.com/wojnosystems/vsql/vrow"
	"github.com/wojnosystems/vsql/vrows"
)

// Count is a convenience method to count the number of results from a vquery
// @param ctx Context to constrain the run-time of this call
// @param vquery the SQL vquery and parameters to use in the call
// @return number how many things are counted
// @return err errors encountered while making the database call. Database-specific errors are likely
func Count(ctx context.Context, queryer vquery.Queryer, q param.Queryer) (number ulong.ULong, err error) {
	var ok bool
	ok, err = vrow.QueryOne(queryer, ctx, q, func(ro vrows.Rower) (err error) {
		return ro.Scan(&number)
	})
	if !ok {
		return 0, sql.ErrNoRows
	}
	return
}
