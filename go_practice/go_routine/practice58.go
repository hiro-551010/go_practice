package main

import "fmt"


func goroutine(words []string, c chan string) {
	sum := ""
	for _, v := range words {
		sum += v
		c <- sum
	}
	close(c)
}


func main() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}
}