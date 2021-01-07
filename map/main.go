package main

import "fmt"

func main() {

	//	var colors map[string]string
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// delete(colors, "white")
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//initalize map
	m := make(map[string]int)

	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is not nil")
	}

	m["Udir"] = 10
	m["Udit2"] = 20
	printM(m)

	printMap(colors)
}

func printM(d map[string]int) {
	for key, value := range d {
		fmt.Println("key is ", key, " value is ", value)
	}
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}
