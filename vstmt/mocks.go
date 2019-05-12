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

package vstmt

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/wojnosystems/vsql/param"
	"github.com/wojnosystems/vsql/vresult"
	"github.com/wojnosystems/vsql/vrows"
)

type StatementQueryerMock struct {
	mock.Mock
}

func (s *StatementQueryerMock) Query(ctx context.Context, query param.Parameterer) (r vrows.Rowser, err error) {
	a := s.Called(ctx, query)
	if a.Get(0) != nil {
		r = a.Get(0).(vrows.Rowser)
	}
	err = a.Error(1)
	return
}

type StatementInserterMock struct {
	mock.Mock
}

func (s *StatementInserterMock) Insert(ctx context.Context, query param.Parameterer) (r vresult.InsertResulter, err error) {
	a := s.Called(ctx, query)
	if a.Get(0) != nil {
		r = a.Get(0).(vresult.InsertResulter)
	}
	err = a.Error(1)
	return
}

type StatementExecerMock struct {
	mock.Mock
}

func (s *StatementExecerMock) Exec(ctx context.Context, query param.Parameterer) (r vresult.Resulter, err error) {
	a := s.Called(ctx, query)
	if a.Get(0) != nil {
		r = a.Get(0).(vresult.Resulter)
	}
	err = a.Error(1)
	return
}

type StatementerMock struct {
	mock.Mock
}

func (s *StatementerMock) Query(ctx context.Context, query param.Parameterer) (r vrows.Rowser, err error) {
	a := s.Called(ctx, query)
	if a.Get(0) != nil {
		r = a.Get(0).(vrows.Rowser)
	}
	err = a.Error(1)
	return
}
func (s *StatementerMock) Insert(ctx context.Context, query param.Parameterer) (r vresult.InsertResulter, err error) {
	a := s.Called(ctx, query)
	if a.Get(0) != nil {
		r = a.Get(0).(vresult.InsertResulter)
	}
	err = a.Error(1)
	return
}
func (s *StatementerMock) Exec(ctx context.Context, query param.Parameterer) (r vresult.Resulter, err error) {
	a := s.Called(ctx, query)
	if a.Get(0) != nil {
		r = a.Get(0).(vresult.Resulter)
	}
	err = a.Error(1)
	return
}
func (s *StatementerMock) Close() (err error) {
	a := s.Called()
	err = a.Error(0)
	return
}

type PreparerMock struct {
	mock.Mock
}

func (s *PreparerMock) Prepare(ctx context.Context, query param.Queryer) (st Statementer, err error) {
	a := s.Called(ctx, query)
	if a.Get(0) != nil {
		st = a.Get(0).(Statementer)
	}
	err = a.Error(1)
	return
}
