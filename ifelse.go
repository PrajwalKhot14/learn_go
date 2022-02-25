package main

func main() {
	if 9%2 == 0 {
		println("Even")
	} else {
		println("Odd")
	}

	var n = 10

	if n < 0 {
		println("Negative")
	} else if n < 9 {
		println("Less than 10")
	} else {
		println("Greater than 10")
	}

}
