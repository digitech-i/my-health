// main.go

package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(3, 5)
	fmt.Println("Result:", result)
}
