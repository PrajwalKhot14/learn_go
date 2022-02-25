package main

func main() {
	i := 1
	for i <= 3 {
		println(i)
		i += 1
	}
	println()
	for j := 0; j <= 3; j++ {
		println(j)
	}
	println()

	for {
		println("Break")
		break
	}
	println()
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		println(n)
	}
}
