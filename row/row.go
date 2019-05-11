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
	"vsql/param"
	"vsql/query"
	"vsql/rows"
)

// Each is a convenience method to help take are of converting multiple rows of similar types into your custom types.
//
// Each will also handle cleaning up the Rowser record when it returns, preventing memory leaks
//
// @param r comes from Query() calls
// @param eachRow is a predicate to call for each row encountered. Return false, nil to keep going. Return true, nil to stop the loop and clean up the row. If you return an error, that error will be passed to the caller of Each and iteration will stop as well.
// @return err the database error encountered or that was returned from eachRow
//
// this will not return that annoying sql.ErrNoRows error. Usually, you're building an array with this and NoRows is not an error, but a valid state.
//
// Example:
//   users := make([]User,0,10)
//   q, err := db.Query(queryParam)
//   if err != nil { return err }
//   err := sqlface.Each( q, func(row Rower)(stop bool, err error) {
//     u := User
//     err = row.Scan(&u.name, &u.age)
//     if err == nil {
//       users = append(users, user)
//     }
//     return true, err
//   } )
func Each( r rows.Rowser, eachRow func(ro rows.Rower)(stop bool, err error) ) (err error) {
	defer func() { _ = r.Close() }()
	var ro rows.Rower
	for {
		ro = r.Next()
		if ro == nil {
			break
		} else {
			var stop bool
			stop, err = eachRow(ro)
			if stop || err != nil {
				// eachRow returned an error or has indicated that iteration should not continue, but not return an error
				return
			}
		}
	}
	return
}

func QueryEach( queryer query.Queryer, ctx context.Context, q param.Queryer, eachRow func(ro rows.Rower)(cont bool, err error) ) (err error) {
	var qr rows.Rowser
	qr, err = queryer.Query(ctx, q)
	if err == sql.ErrNoRows {
		// hide the noRows error, dumb interface decision to use error for this state... :(
		if qr != nil {
			_ = qr.Close()
		}
		return nil
	}
	if err != nil {
		// return the error
		return
	}
	return Each( qr, eachRow )
}

// One is a convenience method to help take are of converting a single row into a single item
//
// One will also handle cleaning up the Rowser record when it returns, preventing memory leaks
//
// @param r comes from Query() calls
// @param theRow is a predicate to call for the top row encountered. If you return an error, that error will be passed to the caller of One
// @return ok true if the database returned at least 1 result, false if nothing was returned
// @return err the database error encountered or that was returned from eachRow
//
// this will not return that annoying sql.ErrNoRows error. Usually, you're building an array with this and NoRows is not an error, but a valid state.
func One( r rows.Rowser, theRow func(ro rows.Rower)(err error)) (ok bool, err error) {
	defer func() { _ = r.Close() }()
	ro := r.Next()
	if ro == nil {
		// nothing returned
		return false, nil
	} else {
		err = theRow(ro)
	}
	return true, err
}

func QueryOne( queryer query.Queryer, ctx context.Context, q param.Queryer, theRow func(ro rows.Rower)(err error) ) (ok bool, err error) {
	var qr rows.Rowser
	qr, err = queryer.Query(ctx, q)
	if err == sql.ErrNoRows {
		// hide the noRows error, dumb interface decision to use error for this state... :(
		if qr != nil {
			_ = qr.Close()
		}
		return false, nil
	}
	if err != nil {
		// return the error
		return
	}
	return One( qr, theRow )
}
