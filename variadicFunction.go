package main

import "fmt"

func sum(nums ...int) int {
	sumval := 0
	for _, v := range nums {
		sumval += v
	}
	return sumval
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	lst := make([]int, 0)
	for i := 1; i < 10; i++ {
		lst = append(lst, i)
	}
	fmt.Println(lst)
	fmt.Println()
	fmt.Println(sum(lst...))
}
