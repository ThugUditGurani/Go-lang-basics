package main

import (
	"fmt"
)

/*
An arrays in a numbered sequence of elements of a specific length
*/

func main() {
	// First Way to define Arrays
	var a [5]int
	fmt.Println("emp: ",a)
	a[4] = 100
	fmt.Println("set: ",a)
	fmt.Println("get:",a[4])
	fmt.Println("len:",len(a))

	//Second way to define Arrays
	b := [4]int{1,2,3,4}
	fmt.Println("Value b:",b)

	// Two Dimensional Arrays
	var twoD [2][3]int
	for i := 0; i< 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d",twoD)

}



