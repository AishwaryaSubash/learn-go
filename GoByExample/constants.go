package main

import (
	"fmt"
	"reflect"
	"math"
)

const str string = "hi"

func constants() {
	const str string = "hi"
	fmt.Println(str)

	// Infers own types
	const d = 3e20 / 500000000
	fmt.Println(d, reflect.TypeOf(d))

	// Explicit type conversion
	const e = int64(d)
	fmt.Println(e, reflect.TypeOf(e))
	
	// Type is assumed when used in a context
	n := math.Sin(500000000)
	fmt.Println(n, reflect.TypeOf(n))
}
