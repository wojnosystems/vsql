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

package aggregator

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/wojnosystems/vsql/ulong"
	"github.com/wojnosystems/vsql/vquery"
	"github.com/wojnosystems/vsql/vrows"
	"testing"
)

func TestCount(t *testing.T) {
	expectCount := ulong.New(5)

	rowMock := &vrows.RowerMock{}
	rowMock.On("Scan", mock.Anything).
		Once().
		Return(nil)
	rowMock.ScanMock = func(values ...interface{}) {
		*values[0].(*ulong.ULong) = expectCount
	}

	rowsMock := &vrows.RowserMock{}
	rowsMock.On("Next", mock.Anything).
		Once().
		Return(rowMock, nil)
	rowsMock.On("Close").
		Once().
		Return(nil)

	queryerMock := &vquery.QueryerMock{}
	queryerMock.On("Query", mock.Anything, mock.Anything).
		Once().
		Return(rowsMock, nil)

	c, err := Count(context.Background(),
		queryerMock, nil)
	if err != nil {
		t.Fatal(err)
	}
	if c != expectCount {
		t.Errorf("Expected count to be %d, but got %d", expectCount, c)
	}

	queryerMock.AssertExpectations(t)
	rowsMock.AssertExpectations(t)
	rowMock.AssertExpectations(t)
}
