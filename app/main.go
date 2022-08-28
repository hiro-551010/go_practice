package main

import (
	"app/mylib"
	"app/mylib/under"
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	fmt.Println(mylib.Average(s))
	mylib.Say()
	under.Ground()

	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)
}
