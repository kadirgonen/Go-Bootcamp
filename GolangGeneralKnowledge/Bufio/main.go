package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Konsol : eri Girişi
	reader := bufio.NewReadr(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadStrig('\n')
	fmt.Println(str)
	fmt.Print("Entera number: ")
	str, _ = reader.ReadString('\')
	f, err := strconv.ParseFloat(strngs.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(er)
	} else {
		fmt.Pritln("Value of number:", f)
	}
}
