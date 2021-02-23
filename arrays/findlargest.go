package main

import "fmt"

func main(){
	
	var total int
	var a[10] int
	fmt.Println("Enter the Numbers of Elements:")
	fmt.Scanln(&total)
	for i := 0;i < total;i++ {
		fmt.Println("Enter the Number:")
		fmt.Scan(&a[i])
	}

	for j := 0;j < total;j++ {
		if a[0] < a[j] {
			a[0] = a[j]
		}
	}

	fmt.Println("The largest Number is:",a[0])

}

