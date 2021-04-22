package main

import "fmt"

type celsiusTwo float64
type fahrenheitTwo float64

func (c celsiusTwo) fahrenheitTwo() fahrenheitTwo {
	return fahrenheitTwo((c * 9.0/5.0) + 32.0)
}
func (f fahrenheitTwo) celsiusTwo() celsiusTwo {
	return celsiusTwo((f -32.0) * 5.0/9.0)
}

const (
	line = "=========================="
	rowFormat = "| %8s | %8s | \n"
	numberFormat = "%.1f"
)

type getRowFn func(row int) (string,string)

func drawTable(hdr1, hdr2 string,rows int, getRow getRowFn)  {
	fmt.Println(line)
	fmt.Println(rowFormat, hdr1, hdr2)
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat,cell1,cell2)
	}
	fmt.Println(line)
}

func ctof(row int) (string,string)  {
	c := celsiusTwo(row*5 - 40)
	f := c.fahrenheitTwo()
	cell1 := fmt.Sprintf(numberFormat, c)
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1,cell2
}

func ftoc(row int) (string, string) {
	f := fahrenheitTwo(row*5 - 40)
	c := f.celsiusTwo()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

func main() {
	drawTable("ºC", "ºF", 29, ctof)
	fmt.Println()
	drawTable("ºF", "ºC", 29, ftoc)
}
