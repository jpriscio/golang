package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	CPF   int
	Ativo bool
}

func (c *Cliente) Desativar() {
	c.Ativo = false
}

func main() {
	cliente1 := Cliente{Nome: "Joao", Idade: 25, CPF: 20845378940, Ativo: true}

	fmt.Printf("Nome: %s\nIdade: %d\nCPF: %d\nSituação no sistema: %t\n", cliente1.Nome, cliente1.Idade, cliente1.CPF, cliente1.Ativo)

	cliente1.Desativar()

	fmt.Printf("Nome: %s\nIdade: %d\nCPF: %d\nSituação no sistema: %t\n", cliente1.Nome, cliente1.Idade, cliente1.CPF, cliente1.Ativo)
}
