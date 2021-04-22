package main

import "fmt"

func main() {

	const distance = 236000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const daysPerYear = 365

	const years = distance / lightSpeed / secondsPerDay / daysPerYear

	fmt.Printf("Canis Major Dwarf Galaxy is ",years, "Light years away.")


}
