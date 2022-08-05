package main

import (
	"fmt"
	"time"
	"os/user"
)

func main() {
	fmt.Println("Hello", time.Now())
	fmt.Println(user.Current())
}