package main

import "fmt"

func init() {
	fmt.Println("init")
}

func buzz() {
	fmt.Println("buzz")
}

func main() {
	buzz()
	fmt.Println("Hello World", "TEST TEST")
}