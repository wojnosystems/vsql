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

package vrow

import (
	"context"
	"database/sql"
	"github.com/wojnosystems/vsql/vparam"
	"github.com/wojnosystems/vsql/vquery"
	"github.com/wojnosystems/vsql/vrows"
)

// Each is a convenience method to help take are of converting multiple vrows of similar types into your custom types.
//
// Each will also handle cleaning up the Rowser record when it returns, preventing memory leaks
//
// @vparam r comes from Query() calls
// @vparam eachRow is a predicate to call for each vrow encountered. Return false, nil to keep going. Return true, nil to stop the loop and clean up the vrow. If you return an error, that error will be passed to the caller of Each and iteration will stop as well.
// @return err the database error encountered or that was returned from eachRow
//
// this will not return that annoying sql.ErrNoRows error. Usually, you're building an array with this and NoRows is not an error, but a valid state.
//
// Example:
//   users := make([]User,0,10)
//   q, err := db.Query(queryParam)
//   if err != nil { return err }
//   err := sqlface.Each( q, func(vrow Rower)(stop bool, err error) {
//     u := User
//     err = vrow.Scan(&u.name, &u.age)
//     if err == nil {
//       users = append(users, user)
//     }
//     return true, err
//   } )
func Each(r vrows.Rowser, eachRow func(ro vrows.Rower) (stop bool, err error)) (err error) {
	defer func() { _ = r.Close() }()
	var ro vrows.Rower
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

func QueryEach(queryer vquery.Queryer, ctx context.Context, q vparam.Queryer, eachRow func(ro vrows.Rower) (cont bool, err error)) (err error) {
	var qr vrows.Rowser
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
	return Each(qr, eachRow)
}

// One is a convenience method to help take are of converting a single vrow into a single item
//
// One will also handle cleaning up the Rowser record when it returns, preventing memory leaks
//
// @vparam r comes from Query() calls
// @vparam theRow is a predicate to call for the top vrow encountered. If you return an error, that error will be passed to the caller of One
// @return ok true if the database returned at least 1 vresult, false if nothing was returned
// @return err the database error encountered or that was returned from eachRow
//
// this will not return that annoying sql.ErrNoRows error. Usually, you're building an array with this and NoRows is not an error, but a valid state.
func One(r vrows.Rowser, theRow func(ro vrows.Rower) (err error)) (ok bool, err error) {
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

func QueryOne(queryer vquery.Queryer, ctx context.Context, q vparam.Queryer, theRow func(ro vrows.Rower) (err error)) (ok bool, err error) {
	var qr vrows.Rowser
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
	return One(qr, theRow)
}
