package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// criando o arquivo

	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	// escrevendo dados nele

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes \n", tamanho)
	f.Close()

	// lendo dados do arquivo

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))
	f.Close()

	// lendo de pouco a pouco

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

}
