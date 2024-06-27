package main

import "fmt"

func slices() {

	// Make a Slice with the params: Type, Length, Capacity
	s := make([]string, 3, 3)
	fmt.Println(s, len(s), cap(s))

	// Direct assignment
	s[0] = "hi"
	fmt.Println(s, len(s), cap(s))

	// Insertion using append method
	s = append(s, "hello", "me")
	fmt.Println(s, len(s), cap(s))

	for i := range len(s) {
		fmt.Println(s[i])
	}
	fmt.Println()

	// Copy a slice
	c := make([]string, 3, 6)
	copy(c, s)
	fmt.Println("copy:", c)

	// Slice operator with slices
	l := s[:4]
	fmt.Println("slice:", l)

	// Multi-dimensional slices can have sub-slices of different sizes
	sub := make([][]int, 3)
	sub[0] = append(sub[0], 0)
	sub[1] = append(sub[1], 1, 2)
	sub[2] = append(sub[2], 3, 4, 5)
	fmt.Println("2D slice:", sub)
}
