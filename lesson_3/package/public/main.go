package main

import (
	"fmt"

	"github.com/kadirgonen/Go-Bootcamp/lesson_3/package/public/formatter"
	"github.com/kadirgonen/Go-Bootcamp/lesson_3/package/public/math"
)

func main() {
	num := math.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}
