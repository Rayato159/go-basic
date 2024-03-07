package main

import "fmt"

func main() {
	numsInt := []int{1, 2, 3}
	numsFloat64 := []float64{1.1, 2.2, 3.3}

	sumInt := sum(numsInt)
	sumFloat64 := sum(numsFloat64)

	fmt.Println(sumInt)
	fmt.Println(sumFloat64)
}

type (
	Number interface {
		int | float64
	}

	Player[T Number] struct {
		Name   string
		Age    int
		Damage T
	}

	Database[T Number] interface {
	}
)

func sum[T Number](nums []T) T {
	var sum T

	for _, n := range nums {
		sum += n
	}

	return sum
}
