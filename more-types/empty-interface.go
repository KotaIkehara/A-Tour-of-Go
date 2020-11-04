package main

import "fmt"

func main() {
	//空のインターフェイスという．
	//任意の方の値を保持できる
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
