//Copyright 2019 Chris Wojno
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package vresult

import (
	"database/sql"
	"github.com/wojnosystems/vsql/ulong"
)

type Resulter interface {
	// RowsAffected is the number of rows created or modified in the previous Exec/Insert call. Not all databases support this
	RowsAffected() (rowsAffected ulong.ULong, err error)
}

type InsertResulter interface {
	// LastInsertId is the ID of the row most recently created by the Insert call. Not all databases support this
	LastInsertId() (id ulong.ULong, err error)
	Resulter
}

// QueryResult is the implementation of the Resulter interface
type QueryResult struct {
	Resulter
	SqlRes sql.Result
}

// RowsAffected is the number of vrows that were altered/created/deleted from the call that produced this Resulter
func (r *QueryResult) RowsAffected() (ulong.ULong, error) {
	v, err := r.SqlRes.RowsAffected()
	if err != nil {
		v = 0
	}
	return ulong.NewInt64(v), err
}

// RowsAffected is ID of the vrow that was inserted from the call that produced this Resulter. Not all databases support this
func (r *QueryResult) LastInsertId() (ulong.ULong, error) {
	v, err := r.SqlRes.LastInsertId()
	if err != nil {
		v = 0
	}
	return ulong.NewInt64(v), err
}
