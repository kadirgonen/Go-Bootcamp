package main

import (
	"fmt"
	
	"github.com/kadirgonen/Go-Bootcamp/lesson_3/package/start_with_lower_case/formatter"
	"github.com/kadirgonen/Go-Bootcamp/lesson_3/package/start_with_lower_case/math"
)

func main() {
	num := math.double(2)
	output := formatter.format(num)
	fmt.Println(output)
}
