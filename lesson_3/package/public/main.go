package main

import (
	"fmt"

	"github.com/kadirgonen/GO-BOOTCAMP/formatter"
	"github.com/kadirgonen/GO-BOOTCAMP/math"
)

func main() {
	num := math.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}
