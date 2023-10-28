package main

import "fmt"

func main() {

	// Array literals
	arr0 := [2]int{1, 2}
	arr1 := arr0 // tham trị

	arr2 := &arr0 // tham chiếu
	arr2[0] = 0

	fmt.Println("arr0", arr0)
	fmt.Println("arr1", arr1) // tham trị
	fmt.Println("arr2", arr2) // tham chiếu
	fmt.Println("")

	// Slice array
	arrA := []int{1, 2, 3, 4, 5, 6, 7}
	arrB := arrA[:] // slice from 0 to end
	arrC := arrA[3:]
	arrD := arrA[:6]
	arrE := arrA[3:6]

	fmt.Println("arrA", arrA) // arrA [1 2 3 4 5 6 7]
	fmt.Println("arrB", arrB) // arrA [1 2 3 4 5 6 7]
	fmt.Println("arrC", arrC) // arrC [4 5 6 7]
	fmt.Println("arrD", arrD) // arrD [1 2 3 4 5 6]
	fmt.Println("arrE", arrE) // arrE [4 5 6]
	fmt.Println("")
	// Array make

	arrX := make([]int, 0)

	arrX = append(arrX, 1)
	arrX = append(arrX, 2)
	arrX = append(arrX, []int{3, 4}...)
	arrX = append(arrX, 5)
	fmt.Println(arrX, len(arrX), cap(arrX))
}
