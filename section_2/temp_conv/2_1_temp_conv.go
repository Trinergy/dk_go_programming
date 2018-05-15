// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

// Celsius is a unit of measuring temperature
type Celsius float64

// Fahrenheit is a unit of measuring temperature
type Fahrenheit float64

// Kelvin is a unit of measuring temperature
type Kelvin float64

const (
	// AbsoluteZeroC is the temperature for absolute 0 in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is the freezing temperature in Celsius
	FreezingC Celsius = 0
	// BoilingC is the boiling temperature in Celsius
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

// KToC converts Kelvin to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF converts Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k*9/5 - 459.67) }

// CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC converts Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) * 5 / 9) }
