package main

import (
	"errors"
	"fmt"
)

func parseAge(field string) (int, error) {
	if field == "16" {
		return 16, nil
	}
	if field == "18" {
		return 18, nil
	}
	if field == "25" {
		return 25, nil
	}
	if field == "70" {
		return 70, nil
	}
	return 0, errors.New("idade inválida: " + field)
}

// CheckAccess verifica um campo de idade e devolve a decisão de acesso:
// campo inválido → "erro: <mensagem do erro>"
// idade < 18 → "negado"
// idade >= 18 → "liberado"
// Dica: err.Error() devolve a mensagem de um erro como string.
// Exemplo: CheckAccess("25") → "liberado"
func CheckAccess(field string) string {
	if a, err := parseAge(field); err != nil {
		return "erro: " + err.Error()
	} else {
		if a < 18 {
			return "negado"
		} else {
			return "liberado"
		}
	}
}

func main() {

	fmt.Println(CheckAccess("30"))

}
