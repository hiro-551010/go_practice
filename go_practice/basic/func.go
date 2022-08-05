package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addSub(x, y int) (int, int) {
	return x + y, x - y
}

func calc(price, item int) (result int) {
	// result := price * item
	result = price * item
	return
}

func main() {
	r := add(1 ,2)
	// r = "aaa"
	fmt.Println(r)

	r2, r3 := addSub(3, 2)
	fmt.Println(r2, r3)

	r4 := calc(100, 2)
	fmt.Println(r4)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)
}