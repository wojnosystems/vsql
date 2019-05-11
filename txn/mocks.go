package txn

import (
	"database/sql"
	"github.com/stretchr/testify/mock"
)

type TxOptionerMock struct {
	mock.Mock
}
func (t TxOptionerMock)IsolationLevel() sql.IsolationLevel {
	a := t.Called()
	return a.Get(0).(sql.IsolationLevel)
}
func (t *TxOptionerMock)SetIsolationLevel( x sql.IsolationLevel ){
	t.Called(x)
}
func (t TxOptionerMock)ReadOnly() bool {
	a := t.Called()
	return a.Bool(0)
}
func (t *TxOptionerMock)SetReadOnly( x bool ){
	t.Called(x)
}
func (t *TxOptionerMock)ToTxOptions() (o *sql.TxOptions) {
	a := t.Called()
	if a.Get(0) != nil {
		o = a.Get(0).(*sql.TxOptions)
	}
	return
}