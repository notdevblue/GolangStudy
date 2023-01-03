package main

import (
	"fmt"
)

func main() {
	const (
		A int    = 10
		B        = "Hello"
		C string = "CString"
	)

	var (
		a string  = "Hello"
		b         = 10
		c float32 = 1.54
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	name := "Han"

	const GENTOO string = "Gentoo"
	const LINUX = "Linux"

	fmt.Println("Name is " + name)
}
