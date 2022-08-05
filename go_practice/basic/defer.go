package main

import (
	"fmt"
	"os"
)

func engage() {
	defer fmt.Println("end")

	fmt.Println("start")
}

func main() {
	// engage()
	// defer fmt.Println("world")
	// engage()
	// fmt.Println("hello")
	// engage()

	// fmt.Println("run")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("success")

	file, _ := os.Open("./switch.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}