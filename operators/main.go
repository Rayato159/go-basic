package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println("Add", a+b)
	fmt.Println("Subtract", a-b)
	fmt.Println("Multiply", a*b)
	fmt.Println("Divide", a/b)
	fmt.Println("Mod", a%b)

	fmt.Println()

	fmt.Println("Is", a == b)
	fmt.Println("Not", a != b)
	fmt.Println("More than", a > b)
	fmt.Println("Less than", a < b)
	fmt.Println("More than or equal to", a >= b)
	fmt.Println("Less than or equal to", a <= b)

	fmt.Println()

	fmt.Println("And", true && true)
	fmt.Println("Or", true || false)
}
