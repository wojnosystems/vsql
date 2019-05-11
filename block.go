package vsql

import (
	"context"
	"fmt"
	"runtime/debug"
	"vsql/vtxn"
)

// Txn creates a transaction in a block, vastly cleaning up transaction boiler-plate
// @param s is the database connection (SQLer interface implementation) to use to start the transaction
// @param ctx is the context to use when starting the transaction
// @param txOps are the options to use when starting the transaction, or nil to use the default
// @param block is the func closure to use within the transaction. When this method ends, the transaction will either be rolled back or committed. If you pass true for rollback or return non-nil for error, the transaction will be rolled back. If rollback is false (the default) and the err is nil (the default), then the transactions will be committed
// @return err the error encountered during Begin, your block call, Rollback, or Commit
func Txn(s SQLer, ctx context.Context, txOps vtxn.TxOptioner, block func(t QueryExecer) (rollback bool, err error)) (err error) {
	var tx QueryExecTransactioner
	tx, err = s.Begin(ctx, txOps)
	if err != nil {
		return
	}
	func() {
		defer func() {
			// This defer ensures that we rollback transactions, even when panics occur
			if r := recover(); r != nil {
				_ = tx.Rollback()
				// regurgitate the panic for debugging
				panic(fmt.Errorf(`panic: %v\n%s`, r, debug.Stack()))
			}
		}()
		rollback := true
		rollback, err = block(tx)
		if rollback || err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	return
}

// TxnNested creates a transaction in a block but also allows for nested transactions, assuming the implementation supports that
// @param s is the database connection (NestedSQLer interface implementation) to use to start the transaction
// @param ctx is the context to use when starting the transaction
// @param txOps are the options to use when starting the transaction, or nil to use the default
// @param block is the func closure to use within the transaction. When this method ends, the transaction will either be rolled back or committed. If you pass true for rollback or return non-nil for error, the transaction will be rolled back. If rollback is false (the default) and the err is nil (the default), then the transactions will be committed
// @return err the error encountered during Begin, your block call, Rollback, or Commit
func TxnNested(s NestedSQLer, ctx context.Context, txOps vtxn.TxOptioner, block func(t QueryExecTransactioner) (rollback bool, err error)) (err error) {
	var tx QueryExecTransactioner
	tx, err = s.Begin(ctx, txOps)
	if err != nil {
		return
	}
	func() {
		defer func() {
			// This defer ensures that we rollback transactions, even when panics occur
			if r := recover(); r != nil {
				_ = tx.Rollback()
				// regurgitate the panic for debugging
				panic(fmt.Errorf(`panic: %v\n%s`, r, debug.Stack()))
			}
		}()
		rollback := true
		rollback, err = block(tx)
		if rollback || err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	return
}
