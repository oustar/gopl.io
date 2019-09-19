package charcount

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	var tests = []struct {
		input string
		want  map[rune]int
	}{
		{"a", map[rune]int{'a': 1}},
		{"世界", map[rune]int{'世': 1, '界': 1}},
	}

	for _, test := range tests {
		got, err := count(strings.NewReader(test.input))
		if err != nil {
			t.Errorf("read error:%s(%v)", test.input, err)
			continue
		}
		for r, n := range test.want {
			if got[r] != n {
				t.Errorf("read:%s, got=%v, want=%v", test.input, got, test.want)
			}
		}
	}
}
