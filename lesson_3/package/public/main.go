package main

import (
	"fmt"

	"github.com/kadirgonen/Go-Bootcamp/formatter"
	"github.com/kadirgonen/Go-Bootcamp/math"
)

func main() {
	num := math.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}
