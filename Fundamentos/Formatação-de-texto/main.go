package main

import "fmt"

func PriceLabel(name string, cents int, onSale bool) string {
	if onSale == true {
		return fmt.Sprintf("PROMO %s: R$ %d.%02d", name, cents/100, cents%100)
	}
	return fmt.Sprintf("%s: R$ %d.%02d", name, cents/100, cents%100)
}

func main() {
	fmt.Println(PriceLabel("Camisa", 49, true))
}
