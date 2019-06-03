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

package interpolation_strategy

// InterpolateStrategy is how to replace the placeholders in the string component of the vquery when a variable needs to be inserted
// Sadly, this is driver/database-dependent. MySQL uses positional question marks (?) while Postgres uses numbered (ordinal) position markers e.g. $1, $5, etc.
// This strategy is how the driver will tell the Parameter interface how to build placeholders.
// The good news is, once the driver has an InterpolationStrategy, then the driver can use these parameter interfaces. The implementation of the easy_sql interface must instantiate and create this strategy as needed. While MySQL's replacement strategy is very simple, just replace with question marks, the Postgres implementation will need to store state and increment the state as InsertPlaceholderIntoSQL is called.
// InsertPlaceholderIntoSQL will only be called once for each placeholder to be inserted into the string and it will be called in order and every value will be used
type InterpolateStrategy interface {
	// InsertPlaceholderIntoSQL returns the string to insert at the position a placeholder needs to appear in the SQL vquery such that the driver is able to replace that with the parameterized values
	InsertPlaceholderIntoSQL() string
}

// InterpolationStrategyFactory builds InterpoateStrategies. Used in database drivers to tell implementing libraries how to do variable injection
// This type is in the VSQL library for convenience
type InterpolationStrategyFactory func() InterpolateStrategy
