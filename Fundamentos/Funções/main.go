package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(soma(0, 0))

}

func soma(a, b int) (int, error) {
	if a+b <= 0 {
		return 0, errors.New("numero negativo ou igual a zero")
	} else {
		return a + b, nil
	}
}
