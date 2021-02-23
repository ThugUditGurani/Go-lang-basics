package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year:")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(),10,64)
	fmt.Printf("You typed: %q",2020-input)
}
