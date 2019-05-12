//Copyright 2019 Chris Wojno
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package vsql

import (
	"context"
	"github.com/wojnosystems/vsql/vtxn"
)

type TransactionStarter interface {
	// Begin starts a transaction and then returns a transaction context that is unable to start transactions of its own, but can execute other SQL queries
	Begin(context.Context, vtxn.TxOptioner) (qet QueryExecTransactioner, err error)
}

type Transactioner interface {
	// Commit ends a transaction by persisting the requested changes
	Commit() error
	// Rollback ends a transaction by not persisting the changes made via queries while within the transaction
	Rollback() error
}

// QueryExecTransactioner is a transaction-aware object that can execute SQL instructions.
// You should avoid using this unless you absolutely need to know that you're in a transaction
type QueryExecTransactioner interface {
	Transactioner
	QueryExecer
}

// NestedQueryExecTransactioner is just like Transactioner but these objects can start sub-transactions
// You should avoid using this unless you absolutely need to know that you're in a transaction or you must start sub-transactions
type NestedQueryExecTransactioner interface {
	QueryExecTransactioner
	NestedTransactionStarter
}

type NestedTransactionStarter interface {
	// Begin starts a transaction and then returns a transaction context that is able to start transactions of its own
	Begin(context.Context, vtxn.TxOptioner) (nt NestedQueryExecTransactioner, err error)
}
