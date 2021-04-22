package main

import "fmt"

func main() {

	messageOne := "Hola Estactic Expacial Internacional"

	for _,c := range messageOne{
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c = c - 26
			}
		}else if c >= 'A' && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c",c)
	}



}
