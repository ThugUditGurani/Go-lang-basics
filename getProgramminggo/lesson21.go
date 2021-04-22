package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Name string `json:"name"`
	Lat float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func main() {
	location := []location{
		{Name: "Bradbury Landing" , Lat: -4.456, Long: 123},
		{Name: "Columbia Memorial", Lat: -14.4234, Long: 175.234},
		{Name: "Chanllenger Memorial", Lat: -1.324, Long: 354.234},
	}

	bytes , error := json.MarshalIndent(location, "", "")
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}
