package main

import (
	"fmt"
	"math/rand"
)

type Kelvin float64
type sensor func() Kelvin

func fakeSensor()  Kelvin{
	return Kelvin(rand.Intn(151) + 250)
}
func calibrate(s sensor,offset Kelvin)  sensor {
	return func() Kelvin {
		return s() + offset
	}
}
func main() {
	var offset Kelvin = 5
	sensor := calibrate(fakeSensor,offset)
	for count:=0;count<10;count++ {
		fmt.Println(sensor())
	}
}
