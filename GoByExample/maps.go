package main

import (
	"fmt"
)

func maps() {

	// Declaration-cum-initialisation
	m := map[string]int{
		"k1": 1,
		"k2": 2,
	}
	fmt.Println(m)

	// Declaration and initialisation
	n := make(map[int]string)
	n[1] = "hello"
	n[2] = "hi"
	fmt.Println(n)

	// Get value using key
	// If key does not exist, Zero-value of value type is returned
	val, isPresent := n[1]
	fmt.Println("Key exists?", isPresent, "\nGet val:", val)

	// Removes a key-value pair from map
	delete(m, "k2")
	fmt.Println("delete:", m)

	// Removes all key-value pairs from map
	clear(m)
	fmt.Println("clear:", m)

	// "maps" package contains utility functions for maps
	map1 := map[int]string{1: "apple", 2: "banana"}
    fmt.Println("map:", map1)

    map2 := map[int]string{1: "apple", 2: "banana"}
	fmt.Println("map:", map2)
    // if maps.Equal(map1, map2) {
    //     fmt.Println("map1 == map2")
    // }
}
