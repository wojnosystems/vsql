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
	"github.com/stretchr/testify/mock"
	"vsql/param"
	"vsql/vresult"
	"vsql/vrows"
	"vsql/vstmt"
	"vsql/vtxn"
)

type PingerMock struct {
	mock.Mock
}

func (m *PingerMock) Ping(ctx context.Context) error {
	a := m.Called(ctx)
	return a.Error(0)
}

type QueryExecerMock struct {
	mock.Mock
}

func (m *QueryExecerMock) Query(ctx context.Context, q param.Queryer) (vrows.Rowser, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vrows.Rowser), a.Error(1)
}
func (m *QueryExecerMock) Insert(ctx context.Context, q param.Queryer) (vresult.InsertResulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.InsertResulter), a.Error(1)
}
func (m *QueryExecerMock) Exec(ctx context.Context, q param.Queryer) (vresult.Resulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.Resulter), a.Error(1)
}

type SQLerMock struct {
	mock.Mock
}

func (m *SQLerMock) Begin(ctx context.Context, tx vtxn.TxOptioner) (QueryExecTransactioner, error) {
	a := m.Called(ctx, tx)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(QueryExecTransactioner), a.Error(1)
}
func (m *SQLerMock) Prepare(ctx context.Context, q param.Queryer) (vstmt.Statementer, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vstmt.Statementer), a.Error(1)
}
func (m *SQLerMock) Ping(ctx context.Context) error {
	a := m.Called(ctx)
	return a.Error(0)
}
func (m *SQLerMock) Query(ctx context.Context, q param.Queryer) (vrows.Rowser, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vrows.Rowser), a.Error(1)
}
func (m *SQLerMock) Insert(ctx context.Context, q param.Queryer) (vresult.InsertResulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.InsertResulter), a.Error(1)
}
func (m *SQLerMock) Exec(ctx context.Context, q param.Queryer) (vresult.Resulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.Resulter), a.Error(1)
}

type NestedSQLerMock struct {
	mock.Mock
}

func (m *NestedSQLerMock) Begin(ctx context.Context, tx vtxn.TxOptioner) (NestedTransactioner, error) {
	a := m.Called(ctx, tx)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(NestedTransactioner), a.Error(1)
}
func (m *NestedSQLerMock) Prepare(ctx context.Context, q param.Queryer) (vstmt.Statementer, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vstmt.Statementer), a.Error(1)
}
func (m *NestedSQLerMock) Ping(ctx context.Context) error {
	a := m.Called(ctx)
	return a.Error(0)
}
func (m *NestedSQLerMock) Query(ctx context.Context, q param.Queryer) (vrows.Rowser, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vrows.Rowser), a.Error(1)
}
func (m *NestedSQLerMock) Insert(ctx context.Context, q param.Queryer) (vresult.InsertResulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.InsertResulter), a.Error(1)
}
func (m *NestedSQLerMock) Exec(ctx context.Context, q param.Queryer) (vresult.Resulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.Resulter), a.Error(1)
}

type PreparerMock struct {
	mock.Mock
}

func (m *PreparerMock) Prepare(ctx context.Context, q param.Queryer) (vstmt.Statementer, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vstmt.Statementer), a.Error(1)
}

type TransactionStarterMock struct {
	mock.Mock
}

func (m *TransactionStarterMock) Begin(ctx context.Context, tx vtxn.TxOptioner) (QueryExecTransactioner, error) {
	a := m.Called(ctx, tx)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(QueryExecTransactioner), a.Error(1)
}

type NestedTransactionStarterMock struct {
	mock.Mock
}

func (m *NestedTransactionStarterMock) Begin(ctx context.Context, tx vtxn.TxOptioner) (NestedTransactioner, error) {
	a := m.Called(ctx, tx)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(NestedTransactioner), a.Error(1)
}
