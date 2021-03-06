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

package vsql

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/wojnosystems/vsql/vparam"
	"github.com/wojnosystems/vsql/vresult"
	"github.com/wojnosystems/vsql/vrows"
	"github.com/wojnosystems/vsql/vstmt"
	"github.com/wojnosystems/vsql/vtxn"
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

func (m *QueryExecerMock) Query(ctx context.Context, q vparam.Queryer) (vrows.Rowser, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vrows.Rowser), a.Error(1)
}
func (m *QueryExecerMock) Insert(ctx context.Context, q vparam.Queryer) (vresult.InsertResulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.InsertResulter), a.Error(1)
}
func (m *QueryExecerMock) Exec(ctx context.Context, q vparam.Queryer) (vresult.Resulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.Resulter), a.Error(1)
}

func (m *QueryExecerMock) Begin(ctx context.Context, tx vtxn.TxOptioner) (QueryExecNestedTransactioner, error) {
	a := m.Called(ctx, tx)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(QueryExecNestedTransactioner), a.Error(1)
}

func (m *QueryExecerMock) Prepare(ctx context.Context, q vparam.Queryer) (vstmt.Statementer, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vstmt.Statementer), a.Error(1)
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
func (m *SQLerMock) Prepare(ctx context.Context, q vparam.Queryer) (vstmt.Statementer, error) {
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
func (m *SQLerMock) Close() error {
	a := m.Called()
	return a.Error(0)
}
func (m *SQLerMock) Query(ctx context.Context, q vparam.Queryer) (vrows.Rowser, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vrows.Rowser), a.Error(1)
}
func (m *SQLerMock) Insert(ctx context.Context, q vparam.Queryer) (vresult.InsertResulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.InsertResulter), a.Error(1)
}
func (m *SQLerMock) Exec(ctx context.Context, q vparam.Queryer) (vresult.Resulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.Resulter), a.Error(1)
}

type SQLNesterMock struct {
	mock.Mock
}

func (m *SQLNesterMock) Begin(ctx context.Context, tx vtxn.TxOptioner) (QueryExecNestedTransactioner, error) {
	a := m.Called(ctx, tx)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(QueryExecNestedTransactioner), a.Error(1)
}
func (m *SQLNesterMock) Prepare(ctx context.Context, q vparam.Queryer) (vstmt.Statementer, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vstmt.Statementer), a.Error(1)
}
func (m *SQLNesterMock) Ping(ctx context.Context) error {
	a := m.Called(ctx)
	return a.Error(0)
}
func (m *SQLNesterMock) Close() error {
	a := m.Called()
	return a.Error(0)
}
func (m *SQLNesterMock) Query(ctx context.Context, q vparam.Queryer) (vrows.Rowser, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vrows.Rowser), a.Error(1)
}
func (m *SQLNesterMock) Insert(ctx context.Context, q vparam.Queryer) (vresult.InsertResulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.InsertResulter), a.Error(1)
}
func (m *SQLNesterMock) Exec(ctx context.Context, q vparam.Queryer) (vresult.Resulter, error) {
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

func (m *PreparerMock) Prepare(ctx context.Context, q vparam.Queryer) (vstmt.Statementer, error) {
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

type TransactionNestedStarterMock struct {
	mock.Mock
}

func (m *TransactionNestedStarterMock) Begin(ctx context.Context, tx vtxn.TxOptioner) (QueryExecNestedTransactioner, error) {
	a := m.Called(ctx, tx)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(QueryExecNestedTransactioner), a.Error(1)
}

type QueryExecTransactionerMock struct {
	QueryExecerMock
}

func (m *QueryExecTransactionerMock) Commit() error {
	a := m.Called()
	return a.Error(0)
}

func (m *QueryExecTransactionerMock) Rollback() error {
	a := m.Called()
	return a.Error(0)
}

func (m *QueryExecTransactionerMock) Prepare(ctx context.Context, q vparam.Queryer) (vstmt.Statementer, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vstmt.Statementer), a.Error(1)
}

type QueryExecNestedTransactionerMock struct {
	QueryExecerMock
}

func (m *QueryExecNestedTransactionerMock) Commit() error {
	a := m.Called()
	return a.Error(0)
}

func (m *QueryExecNestedTransactionerMock) Rollback() error {
	a := m.Called()
	return a.Error(0)
}

func (m *QueryExecNestedTransactionerMock) Begin(ctx context.Context, tx vtxn.TxOptioner) (QueryExecNestedTransactioner, error) {
	a := m.Called(ctx, tx)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(QueryExecNestedTransactioner), a.Error(1)
}

func (m *QueryExecNestedTransactionerMock) Prepare(ctx context.Context, q vparam.Queryer) (vstmt.Statementer, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vstmt.Statementer), a.Error(1)
}
