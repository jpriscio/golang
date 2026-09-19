package main

import (
	"fmt"
)

func main() {

	//Copia simples

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 10)

	copy(slice2, slice1)
	fmt.Println(slice2)

}
