package main

import (
	"fmt"
)

func main() {
	num := math.double(2)
	output := formatter.format(num)
	fmt.Println(output)
}
