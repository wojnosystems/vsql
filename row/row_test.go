//Copyright 2019 Chris Wojno
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package row

import (
	"context"
	"database/sql"
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
	"vsql/query"
	"vsql/rows"
)

func TestEachRow_ErrorEachRow(t *testing.T) {
	// Setup
	errExpected := errors.New("failed")
	rowsMock := &rows.RowserMock{}
	rowsMock.On("Next").
		Once().
		Return(&rows.RowerMock{})
	rowsMock.On("Close").
		Once().
		Return(nil)

	// Perform
	err := Each(rowsMock, func(ro rows.Rower) (stop bool, err error) {
		return false, errExpected
	})

	// Assert
	if err != errExpected {
		t.Error("expected an error to be passed through")
	}
	rowsMock.AssertExpectations(t)
}

func TestEachRow_EachRowStopCont(t *testing.T) {
	// Setup
	times := 3
	rowsMock := &rows.RowserMock{}
	rowsMock.On("Next").
		Times(times).
		Return(&rows.RowerMock{})
	rowsMock.On("Close").
		Once().
		Return(nil)
	i := 0

	// Perform
	err := Each(rowsMock,  func(ro rows.Rower) (stop bool, err error) {
		i++
		return i >= times, nil
	})

	// Assert
	if err != nil {
		t.Error("got error, but was not expecting one")
	}
	rowsMock.AssertExpectations(t)
}

func TestQueryEach_NoRows(t *testing.T) {
	// Setup
	errExpected := sql.ErrNoRows
	qqMock := &query.QueryerMock{}
	qqMock.On("Query", context.Background(), nil ).
		Once().
		Return(nil, errExpected)

	// Perform
	err := QueryEach(qqMock, context.Background(), nil, func(ro rows.Rower) (stop bool, err error) {
		return false, nil
	})

	// Assert
	if err != nil {
		t.Error("expected sql.ErrNoRows to not be passed through")
	}
	qqMock.AssertExpectations(t)
}

func TestQueryEach_EachRowStopCont(t *testing.T) {
	// Setup
	times := 3
	rowsMock := &rows.RowserMock{}
	rowsMock.On("Next").
		Times(times).
		Return(&rows.RowerMock{})
	rowsMock.On("Close").
		Once().
		Return(nil)
	qqMock := &query.QueryerMock{}
	qqMock.On("Query", context.Background(), nil ).
		Once().
		Return(rowsMock, nil)
	i := 0

	// Perform
	err := Each(rowsMock, func(ro rows.Rower) (stop bool, err error) {
		i++
		return i >= times, nil
	})

	// Assert
	if err != nil {
		t.Error("got error, but was not expecting one")
	}
	rowsMock.AssertExpectations(t)
}


func TestOneRow_NoResults(t *testing.T) {
	// Setup
	rowsMock := &rows.RowserMock{}
	rowsMock.On("Next").
		Once().
		Return(nil)
	rowsMock.On("Close").
		Once().
		Return(nil)

	// Perform
	ok, err := One(rowsMock, func(ro rows.Rower) (err error) {
		t.Error("Method should not be called")
		return nil
	})

	// Assert
	if err != nil {
		t.Error("expected no error to be passed through")
	}
	if ok {
		t.Error("expected no rows, should not be OK")
	}
	rowsMock.AssertExpectations(t)
}


func TestQueryOne_NoResults(t *testing.T) {
	// Setup
	errExpected := sql.ErrNoRows
	qqMock := &query.QueryerMock{}
	qqMock.On("Query", context.Background(), nil ).
		Once().
		Return(nil, errExpected)

	// Perform
	ok, err := QueryOne(qqMock, context.Background(), nil, func(ro rows.Rower) (err error) {
		t.Error("Method should not be called")
		return nil
	})

	// Assert
	if err != nil {
		t.Error("expected no error to be passed through")
	}
	if ok {
		t.Error("expected no rows, should not be OK")
	}
	qqMock.AssertExpectations(t)
}

func TestQueryOne_Result(t *testing.T) {
	// Setup
	rowMock := &rows.RowerMock{}
	rowMock.On("Scan", mock.Anything).
		Once().
		Return(nil)
	rowsMock := &rows.RowserMock{}
	rowsMock.On("Next").
		Once().
		Return(rowMock)
	rowsMock.On("Close").
		Once().
		Return(nil)
	qqMock := &query.QueryerMock{}
	qqMock.On("Query", context.Background(), nil ).
		Once().
		Return(rowsMock, nil)

	// Perform
	ok, err := QueryOne(qqMock, context.Background(), nil, func(ro rows.Rower) (err error) {
		var thing string
		return ro.Scan(&thing)
	})

	// Assert
	if err != nil {
		t.Error("expected no error to be passed through")
	}
	if !ok {
		t.Error("expected rows, should be OK")
	}
	qqMock.AssertExpectations(t)
	rowsMock.AssertExpectations(t)
	rowMock.AssertExpectations(t)
}