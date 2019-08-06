// Package unitconv performs Celsius and Fahrenheit conversions.
package unitconv

import "fmt"

// Feet is the lenght unit
type Feet float64

// Meter is the lenght unit
type Meter float64

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// FtToM convert Feet to Meter
func FtToM(f Feet) Meter { return Meter(f / 3.2808) }

// MToFt convert Meter to Feet
func MToFt(m Meter) Feet { return Feet(m * 3.2808) }
