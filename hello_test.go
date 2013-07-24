package main

import (
	"testing"
	"fmt"
)

func TestPrint(t *testing.T) {
	if err := Print("123"); err == nil {
		fmt.Println("returned nil")
	} else {
		fmt.Println("returned err", err)
	}
}
