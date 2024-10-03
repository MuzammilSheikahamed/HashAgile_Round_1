package main

import (
	"fmt"
)

func main() {
	s_input := "abcdef"
	key := 2
	num := len(s_input)
	key = key % num

	rotated := ""

	for i := key; i < num; i++ {
		rotated += string(s_input[i])
	}

	for i := 0; i < key; i++ {
		rotated += string(s_input[i])
	}

	fmt.Println(rotated)
}
