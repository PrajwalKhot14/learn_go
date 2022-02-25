package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "d"
	fmt.Println(s)
	s = append(s, "c")
	fmt.Println(s)
	s2 := make([]string, len(s))
	copy(s2, s)
	fmt.Println(s2)
}
