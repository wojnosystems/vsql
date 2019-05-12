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
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTxn_Commit(t *testing.T) {
	ctx := context.Background()

	qet := &QueryExecTransactionerMock{}
	qet.On("Commit").
		Once().
		Return(nil)

	sqlerMock := &SQLerMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	err := Txn(sqlerMock, ctx, nil, func(t QueryExecer) (commit bool, err error) {
		return true, nil
	})

	if err != nil {
		t.Error("error should not have been returned but got", err)
	}

	sqlerMock.AssertExpectations(t)
	qet.AssertExpectations(t)
}

func TestTxn_Rollback(t *testing.T) {
	ctx := context.Background()

	qet := &QueryExecTransactionerMock{}
	qet.On("Rollback").
		Once().
		Return(nil)

	sqlerMock := &SQLerMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	err := Txn(sqlerMock, ctx, nil, func(t QueryExecer) (commit bool, err error) {
		return
	})

	if err != nil {
		t.Error("error should not have been returned but got", err)
	}

	sqlerMock.AssertExpectations(t)
	qet.AssertExpectations(t)
}

func TestTxn_ErrRollback(t *testing.T) {
	forceErr := errors.New("boom")
	ctx := context.Background()

	qet := &QueryExecTransactionerMock{}
	qet.On("Rollback").
		Once().
		Return(nil)

	sqlerMock := &SQLerMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	err := Txn(sqlerMock, ctx, nil, func(t QueryExecer) (commit bool, err error) {
		return true, forceErr
	})

	if err != forceErr {
		t.Error("error should have been returned but got", err)
	}

	sqlerMock.AssertExpectations(t)
	qet.AssertExpectations(t)
}

func TestTxn_PanicRollback(t *testing.T) {
	ctx := context.Background()

	qet := &QueryExecTransactionerMock{}
	qet.On("Rollback").
		Once().
		Return(nil)

	sqlerMock := &SQLerMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	assert.Panics(t, func() {
		_ = Txn(sqlerMock, ctx, nil, func(t QueryExecer) (commit bool, err error) {
			panic("boom")
			return true, nil
		})
	})

	sqlerMock.AssertExpectations(t)
	qet.AssertExpectations(t)
}

func TestTxnNested_Commit(t *testing.T) {
	ctx := context.Background()

	qet := &QueryExecNestedTransactionerMock{}
	qet.On("Commit").
		Once().
		Return(nil)

	sqlerMock := &SQLNesterMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	err := TxnNested(sqlerMock, ctx, nil, func(t QueryExecTransactioner) (commit bool, err error) {
		return true, nil
	})

	if err != nil {
		t.Error("error should not have been returned but got", err)
	}

	sqlerMock.AssertExpectations(t)
	qet.AssertExpectations(t)
}

func TestTxnNested_Rollback(t *testing.T) {
	ctx := context.Background()

	qet := &QueryExecTransactionerMock{}
	qet.On("Rollback").
		Once().
		Return(nil)

	sqlerMock := &SQLNesterMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	err := TxnNested(sqlerMock, ctx, nil, func(t QueryExecTransactioner) (commit bool, err error) {
		return false, nil
	})

	if err != nil {
		t.Error("error should not have been returned but got", err)
	}

	sqlerMock.AssertExpectations(t)
	qet.AssertExpectations(t)
}

func TestTxnNested_ErrRollback(t *testing.T) {
	forceErr := errors.New("boom")
	ctx := context.Background()

	qet := &QueryExecTransactionerMock{}
	qet.On("Rollback").
		Once().
		Return(nil)

	sqlerMock := &SQLNesterMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	err := TxnNested(sqlerMock, ctx, nil, func(t QueryExecTransactioner) (commit bool, err error) {
		return true, forceErr
	})

	if err != forceErr {
		t.Error("error should have been returned but got", err)
	}

	sqlerMock.AssertExpectations(t)
	qet.AssertExpectations(t)
}

func TestTxnNested_PanicRollback(t *testing.T) {
	ctx := context.Background()

	qet := &QueryExecTransactionerMock{}
	qet.On("Rollback").
		Once().
		Return(nil)

	sqlerMock := &SQLNesterMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	assert.Panics(t, func() {
		_ = TxnNested(sqlerMock, ctx, nil, func(t QueryExecTransactioner) (commit bool, err error) {
			panic("boom")
			return true, nil
		})
	})

	sqlerMock.AssertExpectations(t)
	qet.AssertExpectations(t)
}
