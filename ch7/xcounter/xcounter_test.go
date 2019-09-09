package xcounter

import (
	"bytes"
	"testing"
)

func TestWordCounter(t *testing.T) {
	var tests = []struct {
		s string
		c WordCounter
	}{
		{"", WordCounter(0)},
		{"h", WordCounter(1)},
		{"hello world", WordCounter(2)},
		{"hello world \n hello world", WordCounter(4)},
	}

	for _, test := range tests {
		var c WordCounter
		c.Write([]byte(test.s))
		if c != test.c {
			t.Errorf("WordCounter is err got %d want %d", c, test.c)
		}
	}
}

func TestCountWriter(t *testing.T) {
	var buf bytes.Buffer
	var cw CountWriter
	var pn = cw.CountingWriter(&buf)
	cw.Write([]byte("hello world"))
	if *pn != 11 {
		t.Errorf("Counting Writer is err got %d want 11", *pn)
	}
}
