// helloworld.go
package main

import "fmt"

// var variableName Type

// Variables in block
var (
	gb = "global_block"

	isRs = true
)

// Constant in block
const (
	WIDTH = 20
)

func main() {

	// Basic
	var str string = "Hello Golang"
	fmt.Println(str)

	y := "Let's go go"
	fmt.Println(y)

	// Tham trị
	var i int
	i = 99
	fmt.Println("i =", i)

	j := i
	fmt.Println("j =", j)

	// Tham chiếu

	// Constant
	const LENGTH int = 10
	fmt.Println("Const", LENGTH, " x ", WIDTH)

	// Print the variables in block
	fmt.Println(gb, "\n")

}
