package main

import "fmt"

func main() {
	// var a [3]int = [3]int{1, 2, 3}

	// for _, v := range a {
	// 	fmt.Println(v)
	// }

	// b := []int{1, 2, 3}

	// b := make([]int, 0)

	// b = append(b, 1, 2, 3)

	// fmt.Println(b)

	// b[2] = 5

	// b = append(b, 4)

	// fmt.Println(b)

	a := []int{1, 2, 3}

	b := double(a)

	fmt.Println(a)
	fmt.Println(b)
}

func double(nums []int) []int {
	newNums := make([]int, len(nums))

	for i := range nums {
		newNums[i] = nums[i] * 2
	}

	return newNums
}
