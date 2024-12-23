package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func hello(x int) string {
	if x == 13 {
		return ""
	}
	return "hello"
}
