package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println(m)

	delete(m, "k2")
	fmt.Println(m)

	_, v := m["k1"]
	fmt.Println(v)

	n := map[string]int{"K1": 1, "K2": 2}
	fmt.Println(n)
}
