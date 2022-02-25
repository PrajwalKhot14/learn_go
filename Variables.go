package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = "Hello"
	var b = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))

	var flag = true
	fmt.Println(flag)
}
