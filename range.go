package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for i, num := range nums {
		sum += num
		fmt.Println(i)   //index
		fmt.Println(num) //value
		fmt.Println()
	}
	fmt.Println(sum)

	fruits := map[string]string{"a": "Apple", "b": "Ball"}
	for k, v := range fruits {
		fmt.Println(k, " for ", v)
	}
}
