package main

import "fmt"

func foo() {
	fmt.Println("foo")
	float := 1.2
	var f32 float32 = 2.4
	fmt.Printf("%T\n", float)
	fmt.Printf("%T", f32)
}

func main() {
	// var i int = 1
	// var f64 float64 = 1.2
	// var aaa string = "test"
	// var t bool = true
	// var f bool = false
	var t, f bool = true, false
	var (
		i int = 1
		f64 float64 = 1.2
		aaa string = "test"
	)
	fmt.Println(i, f64, aaa, t, f)
	foo()
}