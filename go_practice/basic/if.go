package main

import "fmt"

func by2(num int) string {
	if num % 2 == 0 {
		return "ok"
	} else {
		return "ng"
	}
}

func main() {
	result := by2(10)
	if result == "ok" {
		fmt.Println("greet")
	}
	if result2 := by2(10); result2 == "ok" {
		fmt.Println("greet 2")
	}
	// num := 4
	// if num % 2 == 0 {
	// 	fmt.Println("even")
	// } else if num % 3 == 0 {
	// 	fmt.Println("odd")
	// } else {
	// 	fmt.Println("else")
	// }

	x, y:= 10, 10
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}
}