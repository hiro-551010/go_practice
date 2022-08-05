package main

import "fmt"

type Vertex struct {
	// 小文字はprivate 大文字はpublic
	X, Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	// v.Y = 2000
	(*v).Y = 2000
}

func main() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	changeVertex2(&v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	// changeVertex(v2)
	changeVertex2(v2)
	fmt.Println(*v2)
	// v := Vertex{X: 1, Y: 2}
	// fmt.Println(v)
	// fmt.Println(v.X, v.Y)

	// v2 := Vertex{X: 1}
	// fmt.Println(v2)

	// v3 := Vertex{1, 2, "test"}
	// fmt.Println(v3)

	// v4 := Vertex{}
	// fmt.Println(v4)

	// var v5 Vertex
	// fmt.Println(v5)

	// v6 := new(Vertex)
	// fmt.Println(v6)

	// v7 := &Vertex{}
	// fmt.Println(v7)
}