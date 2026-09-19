package main

import "fmt"

func main() {
	const (
		numero1 = iota // 0
		numero2        // 1
		numero3        // 2
		numero4        // 3
		numero5        // 4
		numero6        // 5
	)

	fmt.Println(numero1, numero2, numero3)
}
