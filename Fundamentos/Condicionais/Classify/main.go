package main

import "fmt"

func main() {

	fmt.Println(Classify(18))

}

// Classify classifica uma temperatura em graus Celsius:
// abaixo de 15 (exclusive) → "frio"
// de 15 a 27 (inclusive nos dois extremos) → "agradável"
// acima de 27 → "quente"
// Exemplo: Classify(20) → "agradável"
func Classify(celsius int) string {

	if celsius < 15 {
		return "frio"
	}
	if celsius <= 27 {
		return "agradável"
	} else {
		return "quente"
	}
}
