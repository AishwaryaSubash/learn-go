package main

import "fmt"

func conditionals() {
	// Parentheses around conditions - NOT REQD
	// Braces around body - REQD
	num := 1
	if num < 5 {
		fmt.Println(num, "is smaller than 5")
	}

	// Statement preceding conditionals
	if n := 0; n < 0 {
		fmt.Println(n, "is smaller than 0")
	} else if n == 0 {
		fmt.Println(n, "is equal to 0")
	} else {
		fmt.Println(n, "is greater than 0")
	}

	// Switch statement
	// No need for explicit BREAK statements
	i := 2
	switch i {
	case 1, 2:
		fmt.Println(i, "is 1 or 2")
	case 3:
		fmt.Println(i, "is 3")
	default:
		fmt.Println(i, "is neither 1, 2 nor 3")
	}

	// No expression => If-Else construct
	switch {
	case i < 0:
		fmt.Println(i, "is negative")
	default:
		fmt.Println(i, "is positive")
	}
}
