package main

import "fmt"

func main() {
	//defer: すぐに評価されるが，この関数は呼び出し元の関数がreturnするまで延長される
	defer fmt.Println("world")

	fmt.Println("hello")
}
