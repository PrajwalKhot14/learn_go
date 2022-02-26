package main

import "fmt"

func multi(a int, b int) int {
	return a * b
}

func addval(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(multi(3, 4))
	fmt.Println(addval(3, 4))
}
