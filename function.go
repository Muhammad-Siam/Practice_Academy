package main

import "fmt"

func add(x float32, y float32) float32 {
	return y / x
}

func main() {
	fmt.Println(add(42, 100))
	fmt.Println(add(100, 10000))
}
