// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

// Celsius is a unit of measuring temperature
type Celsius float64

// Fahrenheit is a unit of measuring temperature
type Fahrenheit float64

const (
	// AbsoluteZeroC is the temperature for absolute 0 in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is the freezing temperature in Celsius
	FreezingC Celsius = 0
	// BoilingC is the boiling temperature in Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (c Fahrenheit) String() string { return fmt.Sprintf("%g°F", c) }

// CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
