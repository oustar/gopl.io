package charcount

import (
	"bufio"
	"io"
)

// Count got the number of rune
func count(rd io.Reader) (map[rune]int, error) {
	counts := make(map[rune]int)
	in := bufio.NewReader(rd)

	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		counts[r]++
	}
	return counts, nil
}
