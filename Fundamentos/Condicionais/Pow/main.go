package main

import (
	"fmt"
)

func main() {
	fmt.Println(Pow(3, 4))

}

func Pow(base int, exponent int) int {
	var result = 1
	for i := 1; i <= exponent; i++ {
		result *= base
	}
	return result
}
