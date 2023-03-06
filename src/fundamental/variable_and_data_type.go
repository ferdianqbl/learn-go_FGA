package main

import "fmt"

func main() {
	// variable with type
	fmt.Println("=== Variable with type ===")
	var name string = "Go"
	var age int = -10

	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", name, age)

	// variable without type
	fmt.Println("=== Variable without type ===")
	var name2 = "Go"
	var age2 = 10

	fmt.Printf("Hello, my name is %T and I'm %T years old.\n", name2, age2)

	// variable with short declaration
	fmt.Println("=== Variable with short declaration ===")
	name3 := "Go"
	age3 := 10

	fmt.Printf("Hello, my name is %T and I'm %T years old.\n", name3, age3)

	// underscore variable
	fmt.Println("=== Underscore variable ===")
	var test string
	_ = test // use underscore variable to ignore unused variable

	/*
		Basic data type:
		- bool
		- string
		- int, int8, int16, int32, int64
		- uint, uint8, uint16, uint32, uint64, uintptr
		- byte // alias for uint8
		- rune // alias for int32
		- float32, float64
		- complex64, complex128

		Reference data type:
		- pointer
		- map
		- slice
		- channel
		- function

		Interface data type:
		- interface

		Aggregate data type:
		- array
		- struct
	*/

	fmt.Println(5.0 / 3.0)

	// IOTA ==> auto increment
	fmt.Println("=== IOTA ===")
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3)

	const (
		c4 = iota * 2
		c5 = iota * 2
		c6 = iota * 2
	)

	fmt.Println(c4, c5, c6)
}
