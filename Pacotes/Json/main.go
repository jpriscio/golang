package main

import (
	"encoding/json"
	"fmt"
)

type conta struct {
	Nome  string
	Valor int
}

func main() {

	conta1 := conta{"Joao", 30}

	// transformar em json

	res, err := json.Marshal(conta1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// Descodificar um json

	var contaX conta

	json.Unmarshal(res, &contaX)
	fmt.Println(contaX)

}
