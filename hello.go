package main

import "fmt"

func main() {
fmt.Println("Hello, Bangladesh!")
//variable practise
var name string = "Siam"
var age float32 = 16.5
var school string = "KK Govt Institution" 
var roll int = 04
var fvsub string = "Phsyics"
var jscreuslt float32 = 4.71
var fvlang  string = "Golang"

fmt.Println(name, age, school, roll, fvsub, jscreuslt, fvlang)
//short variable practise
hobby, hobby1, hobby2, hobby3 := "Study", "COding", "Reading","Writing"
fmt.Println(hobby, hobby1, hobby2, hobby3)

var chr1 string = "A"
var chr2 string = "B"
var chr3 string = "C"
var chr4 string = "D"
var chr5 string = "E"
var chr6 string = "F"
var chr7 string = "G"
var chr8 string = "E"
var chr9 string = "H"
var chr10 string = "I"
var chr11 string = "J"
var chr12 string = "K"
var chr13 string = "L"
var chr14 string = "M"
var chr15 string = "N"
var chr16 string = "O"
var chr17 string = "P"
var chr18 string = "Q"
var chr19 string = "R"
var chr20 string = "S"
var chr21 string = "T"
var chr22 string = "U"
var chr23 string = "V"

fmt.Println(chr1, chr2, chr3, chr4, chr5, chr6, chr7, chr8, chr9, chr10, chr11, chr12, chr13, chr14, chr15, chr16, chr17, chr18, chr19, chr20, chr21, chr22, chr23)
//simple data input for practise
	fmt.Println("Enter your name & age: ")
	var name1 string
	var age1 int
	fmt.Scanf("%s %d", &name1, &age1)
	fmt.Printf("Your name is %s & age is %d...\n", name1, age1)
//array for practise
var siam [4]string
siam[0] = "My name is siam"
siam[1] = "I am a single boy"
siam[2] = "I love almighty Allah"
siam[3] = "I am a very good boy"

fmt.Println(siam[0:4])

//boolean type variable for practise

fmt.Println(true && true)
  fmt.Println(true && false)
  fmt.Println(true || true)
  fmt.Println(true || false)
  fmt.Println(!true)


}