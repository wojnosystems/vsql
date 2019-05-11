//Copyright 2019 Chris Wojno
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package query

import (
	"context"
	"github.com/stretchr/testify/mock"
	"vsql/param"
	"vsql/result"
	"vsql/rows"
)

type QueryerMock struct {
	mock.Mock
}

func (m *QueryerMock) Query(ctx context.Context, q param.Queryer) (rows.Rowser, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(rows.Rowser), a.Error(1)
}

type InserterMock struct {
	mock.Mock
}

func (m *InserterMock) Insert(ctx context.Context, q param.Queryer) (result.InsertResulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(result.InsertResulter), a.Error(1)
}

type ExecerMock struct {
	mock.Mock
}

func (m *ExecerMock) Exec(ctx context.Context, q param.Queryer) (result.Resulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(result.Resulter), a.Error(1)
}