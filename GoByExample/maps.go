package main

import "fmt"

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

	val := n[1]
	fmt.Println("get val:", val)

	delete(m, "k2")
    	fmt.Println("delete:", m)

	clear(m)
    	fmt.Println("clear:", m)
}
