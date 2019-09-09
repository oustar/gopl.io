package xcounter

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

// WordCounter is the type of counting word
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("wordCounter wirte error (%v)", err)
	}
	*c += WordCounter(count)
	return count, nil
}

// LineCounter is the type of counting lines
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("lineCounter write error (%v)", err)
	}
	*c += LineCounter(count)
	return count, nil
}

// CountWriter is the Writer that can count bytes
type CountWriter struct {
	w io.Writer
	c int64
}

func (cw *CountWriter) Write(p []byte) (n int, err error) {
	n, err = cw.w.Write(p)
	cw.c += int64(n)
	return n, err
}

// CountingWriter is the factory tool
func (cw *CountWriter) CountingWriter(w io.Writer) *int64 {
	cw.w = w
	return &(cw.c)
}
