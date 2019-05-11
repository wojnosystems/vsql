package param

import "github.com/stretchr/testify/mock"

type QueryerMock struct {
	mock.Mock
}

func (q *QueryerMock) Interpolate(strategy InterpolateStrategy) (query string, params []interface{}, err error) {
	a := q.Called(strategy)
	query = a.Get(0).(string)
	params = a.Get(1).([]interface{})
	err = a.Error(2)
	return
}
func (q *QueryerMock) SQLQuery(strategy InterpolateStrategy) string {
	a := q.Called(strategy)
	return a.Get(0).(string)
}
