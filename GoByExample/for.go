package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func forLoop() {
	// While loop
	i := 0
	for i<5 {
		fmt.Println(i, reflect.TypeOf(i))
		i = i+1
	}
	fmt.Println()
	
	// While true
	for {
		fmt.Println(strconv.Itoa(i) + " hi")
		if i == 10 {
			break
		}
		i++
	}
	fmt.Println()
	
	// For loop
	for i:=0 ; i<5 ; i++ {
		fmt.Println(i, reflect.TypeOf(i))
	}
	fmt.Println()
	
	// for i in range(5)
	for i := range 5 {
		fmt.Println(i, reflect.TypeOf(i))
	}
	fmt.Println()
}
