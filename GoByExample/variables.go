package main

import (
	"fmt"
	"reflect"
)

func variables() {
	// Only declaration - zero-valued
	var i int
	fmt.Println(i)

	// Multiple declarations
	var b, c int = 2, 3
	fmt.Println(b, c)

	// Automatic type inference on intialisation
	var a = true
	fmt.Println(a, reflect.TypeOf(a))

	// Shorthand for declaration-cum-intialisation
	// Only inside functions
	f := "hi"
	fmt.Println(f)
}
