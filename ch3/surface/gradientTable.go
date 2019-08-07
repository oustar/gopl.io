package main

import "github.com/lucasb-eyer/go-colorful"

// GradientTable is color map
type GradientTable []struct {
	Col colorful.Color
	Pos float64
}

// GetInterpolatedColorFor is the
func (table GradientTable) GetInterpolatedColorFor(t float64) colorful.Color {
	for i := 0; i < len(table)-1; i++ {
		c1 := table[i]
		c2 := table[i+1]
		if c1.Pos <= t && t <= c2.Pos {
			// We are in between c1 and c2. Go blend them!
			t := (t - c1.Pos) / (c2.Pos - c1.Pos)
			return c1.Col.BlendRgb(c2.Col, t)
		}
	}

	// Nothing found? Means we're at (or past) the last gradient keypoint.
	return table[len(table)-1].Col
}

// MustParseHex This is a very nice thing Golang forces you to do!
// It is necessary so that we can write out the literal of the colortable below.
func MustParseHex(s string) colorful.Color {
	c, err := colorful.Hex(s)
	if err != nil {
		panic("MustParseHex: " + err.Error())
	}
	return c
}
