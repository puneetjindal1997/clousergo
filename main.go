package main

import "fmt"

// Welcome to the channel go guruji

// Topic closure functions

// company => projects => a, b, c, d => emp1

// A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

func company() func() int {
	a := 0
	// b := 0
	// c := 0
	// d := 0
	return func() int {
		a++
		return a
	}
}

func main() {
	val := company()

	fmt.Println(val()) // 1
	fmt.Println(val()) // 2
	fmt.Println(val()) // 3
	fmt.Println(val()) // 4

	v := company()
	fmt.Println(v())
	fmt.Println(val())
}
