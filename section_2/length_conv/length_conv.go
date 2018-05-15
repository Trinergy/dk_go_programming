// Package lengthconv handles Feet and Meters conversion
package lengthconv

import "fmt"

// Feet is a unit of measurement for length
type Feet float64

// FToM Converts Feet to Meters
func FToM(f Feet) Meters      { return Meters(f * 0.3048) }
func (f Feet) String() string { return fmt.Sprintf("%g ft", f) }

// Meters is a unit of measurement for length
type Meters float64

// MToF Converts Meters to Feet
func MToF(m Meters) Feet        { return Feet(m / 0.3048) }
func (m Meters) String() string { return fmt.Sprintf("%g m", m) }
