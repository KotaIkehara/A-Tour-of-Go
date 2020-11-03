package main

import "fmt"

func main() {
	fmt.Println("counting")

	// multi defer: LIFOで実行される
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
