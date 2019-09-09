// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"bytes"
	"fmt"
)

//!+Check
var h int

func (v Var) String() string {

	return string(v)
}

func (l literal) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%g", l)
	return buf.String()
}

func (u unary) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%c", u.op)
	fmt.Fprintf(&buf, "%v", u.x)
	return buf.String()

}
func (b binary) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "(%v", b.x)
	fmt.Fprintf(&buf, "%s", string(b.op))
	fmt.Fprintf(&buf, "%v)", b.y)
	return buf.String()
}

func (c call) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s(", c.fn)
	for i, arg := range c.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%v", arg)
	}
	fmt.Fprintf(&buf, ")")
	return buf.String()
}

//!-Check
