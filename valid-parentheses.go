package main

import "fmt"

func main() {
	var stack []string

	var s = "()[]{}"
	m := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	flag := 0

	for _, runeValue := range s {
		if string(runeValue) == "(" || string(runeValue) == "[" || string(runeValue) == "{" {
			stack = append(stack, string(runeValue))
		} else {
			if len(stack) != 0 {
				if m[string(runeValue)] != stack[len(stack)-1] {
					fmt.Println("false 1")
					flag = 1
					break
				} else {
					stack = stack[:len(stack)-1]
				}
			} else {
				fmt.Println("false 2")
				flag = 1
				break
			}
		}
	}
	if flag==0 {
		fmt.Println("true")
	}
}
