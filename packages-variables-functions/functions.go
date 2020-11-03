package main

import "fmt"

// type comes the cariable name
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
