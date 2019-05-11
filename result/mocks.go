package result

import (
	"github.com/stretchr/testify/mock"
	"vsql/ulong"
)

type ResulterMock struct {
	mock.Mock
}

func (m *ResulterMock) RowsAffected() (ulong.ULong, error) {
	a := m.Called()
	return a.Get(0).(ulong.ULong), a.Error(1)
}

type InsertResulterMock struct {
	ResulterMock
}

func (m *InsertResulterMock) LastInsertId() (ulong.ULong, error) {
	a := m.Called()
	return a.Get(0).(ulong.ULong), a.Error(1)
}