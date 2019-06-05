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

package vparam

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNamedParameter_Interpolate(t *testing.T) {
	cases := map[string]struct {
		queryIn            string
		parametersIn       map[string]interface{}
		queryExpected      string
		parametersExpected []interface{}
	}{
		"basic select pet-first": {
			queryIn:            "select * from my_table where value1 = :pet and value2 = :age",
			parametersIn:       map[string]interface{}{"age": 5, "pet": "puppy"},
			queryExpected:      "select * from my_table where value1 = ? and value2 = ?",
			parametersExpected: []interface{}{"puppy", 5},
		},
		"basic select age-first": {
			queryIn:            "select * from my_table where value1 = :age and value2 = :pet",
			parametersIn:       map[string]interface{}{"age": 5, "pet": "puppy"},
			queryExpected:      "select * from my_table where value1 = ? and value2 = ?",
			parametersExpected: []interface{}{5, "puppy"},
		},
		"nothing": {
			queryIn:            "select * from my_table",
			parametersIn:       map[string]interface{}{},
			queryExpected:      "select * from my_table",
			parametersExpected: []interface{}{},
		},
	}
	for caseName, c := range cases {
		ap := NewNamed(c.queryIn)
		for k, v := range c.parametersIn {
			ap.Set(k, v)
		}
		actualQuery, actualParams, err := ap.Interpolate(&testStrategyDefault)
		if err != nil {
			t.Errorf(`%s: Not expecting Interpolate to return an error`, caseName)
		}
		if actualQuery != c.queryExpected {
			t.Errorf(`%s: Query Expected: "%s" but got "%s"`, caseName, c.queryExpected, actualQuery)
		}
		for i := range c.parametersExpected {
			if actualParams[i] != c.parametersExpected[i] {
				t.Errorf(`%s: Param[%d] Expected: "%v" but got "%v"`, caseName, i, c.parametersExpected[i], actualParams[i])
			}
		}

		// Test the NewAppendWithData
		apwd := NewNamedWithData(c.queryIn, c.parametersIn)
		apwdActualQuery, apwdActualParams, err := apwd.Interpolate(&testStrategyDefault)
		assert.Equal(t, actualQuery, apwdActualQuery)
		assert.Equal(t, actualParams, apwdActualParams)
	}
}

type testStrategy struct {
}

func (m *testStrategy) InsertPlaceholderIntoSQL() string {
	return "?"
}

var testStrategyDefault = testStrategy{}
