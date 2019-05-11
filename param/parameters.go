//Copyright 2019 Chris Wojno
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package param

// Param package describes the interfaces and implementations for types of parameter representations.
// This is my attempt to unify the methodologies of modelling SQL strings across the various database/sql drivers.
// By defining replacement strategies, drivers can be written to use the base parameter library with any implementation (hopefully) without re-writing the base libraries created herein.

// Queryer is the combination of any Parameterer, such as Namer and Appender (or your own, if you wish) coupled with the ability to produce a SQLQuery string.
// When creating methods that take sql queries, such as Exec, Query, etc, this should be the interface you need to use. The easy_sql library uses this heavily for this purpose.
// This is an interface to allow you to customize your parameter handling as you see fit or to future proof your code. This also means you can pick and choose which approach you'd like to take. You can append values or use named values as you wish.
type Queryer interface {
	Parameterer
}

// InterpolateStrategy is how to replace the placeholders in the string component of the vquery when a variable needs to be inserted
// Sadly, this is driver/database-dependent. MySQL uses positional question marks (?) while Postgres uses numbered (ordinal) position markers e.g. $1, $5, etc.
// This strategy is how the driver will tell the Parameter interface how to build placeholders.
// The good news is, once the driver has an InterpolationStrategy, then the driver can use these parameter interfaces. The implementation of the easy_sql interface must instantiate and create this strategy as needed. While MySQL's replacement strategy is very simple, just replace with question marks, the Postgres implementation will need to store state and increment the state as InsertPlaceholderIntoSQL is called.
// InsertPlaceholderIntoSQL will only be called once for each placeholder to be inserted into the string and it will be called in order and every value will be used
type InterpolateStrategy interface {
	// InsertPlaceholderIntoSQL returns the string to insert at the position a placeholder needs to appear in the SQL vquery such that the driver is able to replace that with the parameterized values
	InsertPlaceholderIntoSQL() string
}

// Parameterer is the basic interface that defines things that generate SQL-string + value arrays that will be passed to the db.Query/db.Exec calls
type Parameterer interface {
	// Interpolate injects the parameters into the provided statement
	// @return vquery the string SQL vquery, with the placeholders inserted as per the interpolation strategy
	// @return params the values to inject
	// @return err any errors interpolating the vquery
	Interpolate(strategy InterpolateStrategy) (query string, params []interface{}, err error)

	// SQLQuery is the sql-vquery, ready to be passed to the underlying driver for Exec/Query, with the proper parameter markers already substituted
	SQLQuery(strategy InterpolateStrategy) string
}

// Appender is a type of Parameterer that is simply a list of parameters stuck into a SQL string
type Appender interface {
	// Adds a parameter to the list of variables to parameterize. Values appended will be passed to Query/Exec in the order you called Append
	Append(value interface{})
	Parameterer
}

// Namer is a type of Parameterer that is a SQL-string with a collection of named keys paired with values
// Queries should be written with :named keys, which are prefixed with a colon (:) and consist only of a-zA-Z0-9_ and must not start with a number
type Namer interface {
	Set(name string, value interface{})
	Parameterer
}
