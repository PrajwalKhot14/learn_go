package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println(len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twod [5][5]int
	for i := 0; i < len(twod); i++ {
		for j := 0; j < len(twod[0]); j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println(twod)
}
