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

package vquery

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/wojnosystems/vsql/param"
	"github.com/wojnosystems/vsql/vresult"
	"github.com/wojnosystems/vsql/vrows"
)

type QueryerMock struct {
	mock.Mock
}

func (m *QueryerMock) Query(ctx context.Context, q param.Queryer) (vrows.Rowser, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vrows.Rowser), a.Error(1)
}

type InserterMock struct {
	mock.Mock
}

func (m *InserterMock) Insert(ctx context.Context, q param.Queryer) (vresult.InsertResulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.InsertResulter), a.Error(1)
}

type ExecerMock struct {
	mock.Mock
}

func (m *ExecerMock) Exec(ctx context.Context, q param.Queryer) (vresult.Resulter, error) {
	a := m.Called(ctx, q)
	r := a.Get(0)
	if r == nil {
		return nil, a.Error(1)
	}
	return r.(vresult.Resulter), a.Error(1)
}
