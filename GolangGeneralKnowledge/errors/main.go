package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	//GO hatalara uygulamanın bize verdiği bir değer olarak bakar.
	// nil başlangıç değeri olmayan ifadelere verilen değerdir.

	result, err := evenNum(11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Girdiğiniz sayı: ", result)
	}

	//***************************************

	//error interface türünden bir typedir.

	result1, err1 := sRoot(-5)
	if err != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}

	//**********************************************

	file, err9 := os.Open("test.txt")
	if err9 != nil {
		fmt.Println(err9)
	} else {
		fmt.Println("Dosyamız", file)
	}

}

func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("HATA: Çift sayı girmediniz")
	}
	return num, nil
}

func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Negatif sayıların karekökü alınamaz")
	}
	return math.Sqrt(num), nil
}

/* package main
import (
	"fmt"
	"math"
)
func main() {
	result := sRoot(-4)
	{
		fmt.Println(result)
	}
}
func sRoot(num float64) float64 {
	return math.Sqrt(num)
} */ //NaN döner,beklenmeyen sonuç bir hatadır.

/*
*************************************************** */
//EXERCISES

/*func main() {
	fmt.Print("Lütfen aldığınız notu giriniz: ")
	grade, _ := getGrade()
	var result string
	if grade >= 50 {
		result = "Geçtin"
	} else {
		result = "Kaldın"
	}
	fmt.Println(result)
}*/

func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return num, nil
}

/*
func main() {

	target := numRand(1, 100)

	fmt.Println("1 ile 100 arasındaki sayıyı bulmaya çalışın")

	reader := bufio.NewReader(os.Stdin)

	for attempts := 0; attempts < 10; attempts++ {
		fmt.Println(10-attempts, "hakkınız kaldı")
		fmt.Print("Lütfen tahmininizi yapınız: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		if num > target {
			fmt.Println("Tahmininiz daha büyük, daha küçük bir sayı giriniz.")
		} else if num < target {
			fmt.Println("Tahmininiz daha küçük, daha büyük bir sayı giriniz.")
		} else {
			fmt.Println("Doğru Tahmin, hedf sayı ", target, " idi ", attempts, " seferde bulundunuz")
			break
		}

	}

}

func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

*/
