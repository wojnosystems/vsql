//Copyright 2019 Chris Wojno
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
