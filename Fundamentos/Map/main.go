package main

import "fmt"

func main() {
	meumap := make(map[string]int)

	//meumap1 := map[string]int{"joao": 10, "joa": 20}

	meumap["joao"] = 300
	meumap["laion"] = 500
	meumap["Lara"] = 800

	for nome, salario := range meumap {
		fmt.Printf("Meu nome é %s e meu salario é %d\n", nome, salario)
	}

}
