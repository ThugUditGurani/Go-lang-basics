package main

import "fmt"

func kelvinToCelsius(k float64) float64  {
	return k - 273.15
}

func celsiusToFahrenheit(c float64) float64  {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64  {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}


func main() {
	fmt.Printf("233k is %.2f C\n", kelvinToCelsius(233))
	fmt.Printf("0 k is %.2f F\n",kelvinToFahrenheit(0))
}
