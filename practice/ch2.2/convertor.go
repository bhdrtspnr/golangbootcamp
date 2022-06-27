package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

type Feet float64
type meters float64

type pounds float64
type kilograms float64

func CelsiusToFahrenheit(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FahrenheitToCelsius(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FeetToMeters(f Feet) meters { return meters(f * 0.3048) }
func MetersToFeet(m meters) Feet { return Feet(m * 3.28084) }

func PoundsToKilograms(p pounds) kilograms { return kilograms(p * 0.453592) }
func KilogramsToPounds(k kilograms) pounds { return pounds(k * 2.20462) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (c Fahrenheit) String() string { return fmt.Sprintf("%g°F", c) }
func (f Feet) String() string       { return fmt.Sprintf("%gft", f) }
func (m meters) String() string     { return fmt.Sprintf("%gmt", m) }
func (p pounds) String() string     { return fmt.Sprintf("%glbs", p) }
func (k kilograms) String() string  { return fmt.Sprintf("%gkgs", k) }

func main() {
	menu()
}

func menu() {
	fmt.Println("Please enter your unit:")
	fmt.Println("1: Celsius")
	fmt.Println("2: Fahrenheit")
	fmt.Println("3: Feet")
	fmt.Println("4: Meters")
	fmt.Println("5: Pounds")
	fmt.Println("6: Kilograms")
	fmt.Println("7: Exit")

	var unit int
	fmt.Scanln(&unit)
	switch unit {
	case 1:
		fmt.Println("Please enter your value:")
		var value Celsius
		fmt.Scanln(&value)
		fmt.Println(value.String())
		fmt.Println(CelsiusToFahrenheit(value))
	case 2:
		fmt.Println("Please enter your value:")
		var value Fahrenheit
		fmt.Scanln(&value)
		fmt.Println(value.String())
		fmt.Println(FahrenheitToCelsius(value))
	case 3:
		fmt.Println("Please enter your value:")
		var value Feet
		fmt.Scanln(&value)
		fmt.Println(value.String())
		fmt.Println(FeetToMeters(value))
	case 4:
		fmt.Println("Please enter your value:")
		var value meters
		fmt.Scanln(&value)
		fmt.Println(value.String())
		fmt.Println(MetersToFeet(value))
	case 5:
		fmt.Println("Please enter your value:")
		var value pounds
		fmt.Scanln(&value)
		fmt.Println(value.String())
		fmt.Println(PoundsToKilograms(value))
	case 6:
		fmt.Println("Please enter your value:")
		var value kilograms
		fmt.Scanln(&value)
		fmt.Println(value.String())
		fmt.Println(KilogramsToPounds(value))
	case 7:
		fmt.Println("Goodbye!")
		return
	default:
		fmt.Println("Invalid unit")
		menu()

	}
}
