//Copyright 2019 Chris Wojno
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package vtxn

import "database/sql"

type TxOptioner interface {
	IsolationLevel() sql.IsolationLevel
	SetIsolationLevel(sql.IsolationLevel)

	ReadOnly() bool
	SetReadOnly(bool)

	ToTxOptions() (o *sql.TxOptions)
}
type TxOption struct {
	TxOptioner
	isolationLevel sql.IsolationLevel
	readOnly       bool
}

func (t TxOption) IsolationLevel() sql.IsolationLevel {
	return t.isolationLevel
}
func (t *TxOption) SetIsolationLevel(x sql.IsolationLevel) {
	t.isolationLevel = x
}
func (t TxOption) ReadOnly() bool {
	return t.readOnly
}
func (t *TxOption) SetReadOnly(x bool) {
	t.readOnly = x
}
func (t TxOption) ToTxOptions() *sql.TxOptions {
	r := &sql.TxOptions{
		ReadOnly:  t.readOnly,
		Isolation: t.isolationLevel,
	}
	return r
}
