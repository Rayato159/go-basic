package main

import "fmt"

var graph = map[int][]int{
	1: {2, 3},
	2: {4, 5},
	3: {6},
}

func main() {
	// a := map[int]int{
	// 	1: 1,
	// 	2: 2,
	// 	3: 3,
	// }

	// if _, ok := a[1]; ok {
	// 	fmt.Println("OK")
	// } else {
	// 	fmt.Println("NOT OK")
	// }

	// a := make(map[int]int)

	// a[1] = 1

	dfs(graph, 1, make(map[int]bool))
}

func dfs(graph map[int][]int, vertex int, visted map[int]bool) {
	visted[vertex] = true
	for _, v := range graph[vertex] {
		fmt.Printf("->%d", v)
		dfs(graph, v, visted)
	}
}
