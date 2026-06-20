package main

import "fmt"

func main() {

	meuSlice := []int{1, 2, 3, 4, 5}

	//for i, v := range meuSlice {
	//	fmt.Printf("O meu indice é %d e o meu valor é %d\n", i, v)
	//}

	fmt.Printf("tenho um slice de cap=%d e de tamanho=%d, %d\n", cap(meuSlice), len(meuSlice), meuSlice)

	fmt.Printf("tenho um slice de cap=%d e de tamanho=%d, %d\n", cap(meuSlice[:0]), len(meuSlice[:0]), meuSlice[:0])
}
