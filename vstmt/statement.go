//Copyright 2019 Chris Wojno
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package vstmt

import (
	"context"
	"io"
	"vsql/param"
	"vsql/vresult"
	"vsql/vrows"
)

// Statements are inherently database-specific. Implementations are located with the driver
type StatementQueryer interface {
	Query(ctx context.Context, query param.Parameterer) (rows vrows.Rowser, err error)
}

type StatementInserter interface {
	Insert(ctx context.Context, query param.Parameterer) (result vresult.InsertResulter, err error)
}

type StatementExecer interface {
	Exec(ctx context.Context, query param.Parameterer) (result vresult.Resulter, err error)
}

type Statementer interface {
	io.Closer
	StatementQueryer
	StatementInserter
	StatementExecer
}

type Preparer interface {
	Prepare(ctx context.Context, query param.Queryer) (stmt Statementer, err error)
}
