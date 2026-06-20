package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	fmt.Printf("O cliente %s andou\n", c.nome)
}

func main() {

	slice := []interface{}{"joao", 100, 3.2}

	fmt.Println(slice)

	Priscio := Cliente{"Priscio"}

	Priscio.andou()

}
