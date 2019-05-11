package vstmt

import (
	"context"
	"github.com/stretchr/testify/mock"
	"vsql/param"
	"vsql/vresult"
	"vsql/vrows"
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
