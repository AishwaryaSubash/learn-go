package main

import "fmt"

func arrays() {

	// By default, arrays are zero-valued
	var arr [5]int
	for i := range len(arr) {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// Assigning values
	arr[4] = 5
	for i := range len(arr) {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// Declaration-cum-intialisation
	b := [5]int{1, 2, 3, 4, 5}
	for i := range len(b) {
		fmt.Print(b[i], " ")
	}

	// Compiler calculates the length on its own
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(c))

	// Two-dimensional array
	d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("rows:", len(d), " cols:", len(d[0]))
}
