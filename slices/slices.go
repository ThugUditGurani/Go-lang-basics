package main

import "fmt"

/*
Slices are key data type in Go, giving a more powerful interface to sequences than arrays

*/
func main() {
	var s[]int
	s = append(s, 10)
	fmt.Println(s)

	nonZero := make([]string,3)
	nonZero[0] = "3"
	nonZero[1] = "4"
	nonZero[2] = "5"
	nonZero = append(nonZero,"8","1","20")
	fmt.Println(nonZero)

	c := make([]string,len(nonZero))
	copy(c,nonZero)
	fmt.Println("Copy c:",c)

	l := nonZero[2:5]
	fmt.Println("Sl1:",l)

	z := nonZero[:5]
	fmt.Println(z)

	k := nonZero[2:]
	fmt.Println(k)

	t := []string{"g","h","i"}
	fmt.Println("dcl: ",t)

	twoD := make([][]int,3)
	for i := 0; i < 3;i++ {
		innerLen := i +1
		twoD[i] = make([]int,innerLen)
		for j := 0;j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

}