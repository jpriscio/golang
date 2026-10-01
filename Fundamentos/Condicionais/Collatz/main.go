package main

import "fmt"

func main() {
	fmt.Println(CollatzSteps(1))
}

func CollatzSteps(n int) int {
	if n == 0 {
		return -1
	}
	var cont int
	for n != 1 {

		if n%2 == 0 {
			fmt.Println(n)
			n = n / 2
		} else {
			fmt.Println(n)
			n = (n * 3) + 1
		}
		cont++

	}
	return cont
}
