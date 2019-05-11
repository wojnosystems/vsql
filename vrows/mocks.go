package vrows

import "github.com/stretchr/testify/mock"

type RowserMock struct {
	mock.Mock
}

func (m *RowserMock) Next() Rower {
	a := m.Called()
	r := a.Get(0)
	if r == nil {
		return nil
	}
	return r.(Rower)
}
func (m *RowserMock) Close() error {
	a := m.Called()
	return a.Error(0)
}

type RowerMock struct {
	mock.Mock
}

func (m *RowerMock) Scan(values ...interface{}) error {
	a := m.Called(values)
	return a.Error(0)
}

func (m *RowerMock) Columns() (columnNames []string) {
	a := m.Called()
	return a.Get(0).([]string)
}
