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

	err := Txn(sqlerMock, ctx, nil, func(t QueryExecer) (rollback bool, err error) {
		return
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

	err := Txn(sqlerMock, ctx, nil, func(t QueryExecer) (rollback bool, err error) {
		return true, nil
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

	err := Txn(sqlerMock, ctx, nil, func(t QueryExecer) (rollback bool, err error) {
		return false, forceErr
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
		_ = Txn(sqlerMock, ctx, nil, func(t QueryExecer) (rollback bool, err error) {
			return false, nil
		})
	})

	sqlerMock.AssertExpectations(t)
	qet.AssertExpectations(t)
}

func TestTxnNested_Commit(t *testing.T) {
	ctx := context.Background()

	qet := &NestedQueryExecTransactionerMock{}
	qet.On("Commit").
		Once().
		Return(nil)

	sqlerMock := &NestedSQLerMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	err := TxnNested(sqlerMock, ctx, nil, func(t QueryExecTransactioner) (rollback bool, err error) {
		return
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

	sqlerMock := &NestedSQLerMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	err := TxnNested(sqlerMock, ctx, nil, func(t QueryExecTransactioner) (rollback bool, err error) {
		return true, nil
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

	sqlerMock := &NestedSQLerMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	err := TxnNested(sqlerMock, ctx, nil, func(t QueryExecTransactioner) (rollback bool, err error) {
		return false, forceErr
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

	sqlerMock := &NestedSQLerMock{}
	sqlerMock.On("Begin", ctx, nil).
		Once().
		Return(qet, nil)

	assert.Panics(t, func() {
		_ = TxnNested(sqlerMock, ctx, nil, func(t QueryExecTransactioner) (rollback bool, err error) {
			return false, nil
		})
	})

	sqlerMock.AssertExpectations(t)
	qet.AssertExpectations(t)
}
