package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a+b
}

func sub( a, b int) int {
	return a-b
}

func mul (a, b int) int {
	return a*b
}

func div (a, b int) int {
	return a/b
}

func calc() {
	var opt int
	for {
		fmt.Println("Options: 1-Add, 2-Sub, 3-Mul, 4-Div")
		fmt.Scan(&opt)
		switch (opt) {
		case 1:
			sum := add(1, 2)
			fmt.Println("Sum: ", sum)
		case 2:
			diff := sub(1, 2)
			fmt.Println("Diff: ", diff)
		case 3:
			prod := mul(1, 2)
			fmt.Println("Prod: ", prod)
		case 4:
			quo := div(1, 2)
			fmt.Println("Quo: ", quo)
		default:
			break
		}
	}
		

}