package main

import "fmt"

type UserNotFount struct {
	Username string
}

func (e *UserNotFount) Error() string {
	return fmt.Sprintf("User not fount: %v", e.Username)
}

func myFunc() error {
	isError := false
	if isError {
		return nil
	}
	return &UserNotFount{Username: "someone"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}