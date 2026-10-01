package main

import "fmt"

func main() {

	score := 0

	switch {
	case score < 0:
		fmt.Println("Tokens ilimitados")
	case score > 10:
		fmt.Println("Tokens: 3")
	default:
		fmt.Println("zero token")
	}

}
