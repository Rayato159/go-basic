package main

import "fmt"

func main() {
	// hello("Rayato159")
	// hello("Josh")

	// result := add(4, 2)

	// fmt.Println(result)
	// fmt.Println(sub(4, 2))
	// fmt.Println(mul(4, 2))
	// fmt.Println(div(4, 2))

	// fmt.Println(div(4, 0))

	// n := 10
	// double(n)
	// fmt.Println(n)

	result := func(a, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(result)
}

// func hello(name string) {
// 	fmt.Println("Hello", name)
// }

// func add(a, b int) int {
// 	return a + b
// }

// func sub(a, b int) int {
// 	return a - b
// }

// func mul(a, b int) int {
// 	return a * b
// }

// func div(a, b int) int {
// 	if b == 0 {
// 		fmt.Println("Error: Division by zero.")
// 		return 0
// 	}
// 	return a / b
// }

// func double(n int) {
// 	n *= 2
// }
