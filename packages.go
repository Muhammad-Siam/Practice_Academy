package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(44))
	fmt.Println("My roll number is", rand.Intn(2))
    {
	fmt.Println("My family members are", rand.Intn(5))
	fmt.Println("My brothers are", rand.Intn(2))
	fmt.Println("My sisters are", rand.Intn(3))
	}
}	