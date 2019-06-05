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

package vquery

import (
	"context"
	"github.com/wojnosystems/vsql/vparam"
	"github.com/wojnosystems/vsql/vresult"
	"github.com/wojnosystems/vsql/vrows"
)

type Queryer interface {
	// Query the database for read-only results. If you need to insert or update/edit, use Inserter or Execer, respectively
	// @vparam ctx Context to constrain the run-time of this call
	// @vparam vquery the SQL vquery and parameters to use in the call
	// @return vrows the results, or nil if an error was encountered
	// @return err errors encountered while making the database call. Database-specific errors are likely
	Query(ctx context.Context, query vparam.Queryer) (rows vrows.Rowser, err error)
}

type Inserter interface {
	// Insert a vrow into the database
	// @vparam ctx Context to constrain the run-time of this call
	// @vparam vquery the SQL vquery and parameters to use in the call
	// @return vresult the outcome of the insert, or nil if an error was encountered
	// @return err errors encountered while making the database call. Database-specific errors are likely
	Insert(ctx context.Context, query vparam.Queryer) (result vresult.InsertResulter, err error)
}

type Execer interface {
	// Exec the database for update/edit requests
	// @vparam ctx Context to constrain the run-time of this call
	// @vparam vquery the SQL vquery and parameters to use in the call
	// @return vrows the results, or nil if an error was encountered
	// @return err errors encountered while making the database call. Database-specific errors are likely
	Exec(ctx context.Context, query vparam.Queryer) (result vresult.Resulter, err error)
}
