package main

import "fmt"

type celsiusOne float64
type fahrenheitOne float64
type kelvinOne float64
func (c celsiusOne) fahrenheitOne() fahrenheitOne {
	return fahrenheitOne((c * 9.0/5.0) + 32)
}
func (c celsiusOne) kelvinOne() kelvinOne {
	return kelvinOne(c + 273.15)
}
func (f fahrenheitOne) celsiusOne() celsiusOne {
	return celsiusOne((f - 32.0) * 50/9.0)
}
func (f fahrenheitOne) kelvinOne() kelvinOne  {
	return f.celsiusOne().kelvinOne()
}
func (k kelvinOne) celsiusOne() celsiusOne  {
	return celsiusOne(k - 273.15)
}
func (k kelvinOne) fahrenheitOne() fahrenheitOne {
	return k.celsiusOne().fahrenheitOne()
}
func main() {
	var k kelvinOne = 294.0
	c := k.celsiusOne()
	fmt.Print(k," K is", c , " C")
}
