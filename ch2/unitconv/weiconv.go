// Package unitconv performs Celsius and Fahrenheit conversions.
package unitconv

import "fmt"

// Pound is the weight unit
type Pound float64

// Kilogram is the weight unit
type Kilogram float64

func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

// LbToKg convert Pound to Kilgoram
func LbToKg(p Pound) Kilogram { return Kilogram(p / 2.2046) }

// KgToLb convert Kilgoram to Pound
func KgToLb(k Kilogram) Pound { return Pound(k * 2.2046) }
