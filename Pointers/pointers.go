package main

import "fmt"

func main() {
	// -- The simplest pointer example --
	// & gives you the memory address (pointer)
	// * gives you the value at that address

	x := 42
	fmt.Println("x =", x)

	// p holds the address of x (p is a pointer to x)
	p := &x
	fmt.Println("&x =", p) // prints memory address
	fmt.Println("*p =", *p) // prints 42 — value AT that address

	// Change x through the pointer
	*p = 100
	fmt.Println("x =", x) // 100! value changed via pointer
}
