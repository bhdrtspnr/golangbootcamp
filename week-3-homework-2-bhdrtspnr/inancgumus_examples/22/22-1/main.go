// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number

	namePhoneMap := make(map[string]string)
	namePhoneMap["Bahadir"] = "5342377252"

	productAvailabilityMap := make(map[string]bool)
	productAvailabilityMap["Banana"] = true

	nameMultiplePhoneMap := make(map[string][]string)
	nameMultiplePhoneMap["Bahadir"] = append(nameMultiplePhoneMap["Bahadir"], "05342377252")
	nameMultiplePhoneMap["Bahadir"] = append(nameMultiplePhoneMap["Bahadir"], "05342377253")

	shoppingBasketCustID := make(map[string]string)
	shoppingBasketCustID["1"] = "1"

	fmt.Println(namePhoneMap)
	fmt.Println(productAvailabilityMap)
	fmt.Println(nameMultiplePhoneMap)
	fmt.Println(shoppingBasketCustID)
	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable

	// #3
	// Key        : Last name
	// Element    : Phone numbers

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
}
