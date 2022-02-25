package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	}

	var week string = time.Now().Weekday().String()
	fmt.Println(week)
	switch week {
	case time.Saturday.String(), time.Sunday.String():
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

}
