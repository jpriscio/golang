package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FizzBuzzLine(7))
}

// FizzBuzzLine devolve a linha do jogo FizzBuzz para um número:
// múltiplo de 3 e de 5 → "FizzBuzz"
// múltiplo só de 3 → "Fizz"
// múltiplo só de 5 → "Buzz"
// caso contrário → o próprio número como texto (use fmt.Sprintf ou strconv).
// Exemplo: FizzBuzzLine(15) → "FizzBuzz"; FizzBuzzLine(7) → "7"
func FizzBuzzLine(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(n)
	}

}
