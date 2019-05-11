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
	rowMock.ScanTransform = func(values ...interface{}) {
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
