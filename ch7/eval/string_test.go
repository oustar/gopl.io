package eval

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {

	expr, err := Parse("pow(x, 3) + pow(y, 3)")
	if err != nil {
		t.Error(err) // parse error
		return
	}
	fmt.Printf("%v", expr)
}
