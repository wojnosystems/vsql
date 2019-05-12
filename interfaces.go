//Copyright 2019 Chris Wojno
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package vsql

import (
	"github.com/wojnosystems/vsql/pinger"
	"github.com/wojnosystems/vsql/vquery"
	"github.com/wojnosystems/vsql/vstmt"
)

// SQLer is what most database references should use to identify a database "connection"
type SQLer interface {
	TransactionStarter
	pinger.Pinger
	QueryExecer
}

// SQLNester is what database references that support nested transactions should use to identify a database "connection"
// It's the nested transaction version of SQLer
type SQLNester interface {
	TransactionNestedStarter
	pinger.Pinger
	QueryExecer
}

// QueryExecer is the interface to use in functions that need to execute queries without knowing the current transaction state.
// If you're writing methods that CRUD objects/resources, use this interface. If you're sure you only need to read, use Queryer, if you need to exec arbitrary queries, use Execer. Use Inserter if you need to insert values and get the LastInsertId
type QueryExecer interface {
	vquery.Queryer
	vquery.Inserter
	vquery.Execer
	vstmt.Preparer
}
