package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	// a[2] = 300

	var b [2]int = [2]int{100, 200}

	var c[]int = []int{100, 200}
	c = append(c, 300)

	fmt.Println(a, b, c)

	d := []int{1, 2, 3, 4, 5}
	fmt.Println(d[2:4])
	fmt.Println(d[:4])

	var board = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board)

	e := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v", len(e), cap(e), e)

	// f := make([]int, 5)
	f := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		f = append(f, i)
		fmt.Println(f)
	}
	fmt.Println(f)
}