package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100

	AbsoluteZeroK Fahrenheit = -459.67
	FreezingK     Fahrenheit = 32
	BoilingK      Fahrenheit = 212
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (c Fahrenheit) String() string { return fmt.Sprintf("%g°F", c) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Println("Celsius\tFahrenheit")
	fmt.Printf("%g\t%g\n", AbsoluteZeroC, CToF(AbsoluteZeroC))
	fmt.Printf("%g\t%g\n", FreezingC, CToF(FreezingC))
	fmt.Printf("%g\t%g\n", BoilingC, CToF(BoilingC))
	fmt.Printf("%g\t%g\n", FreezingK, FToC(FreezingK))
	fmt.Printf("%g\t%g\n", BoilingK, FToC(BoilingK))
}
