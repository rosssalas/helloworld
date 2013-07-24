package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

func Print(s string) error {
	_, err := fmt.Println(s)
	return err
}
