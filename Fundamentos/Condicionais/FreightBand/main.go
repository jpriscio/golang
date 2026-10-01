package main

// FreightBand devolve a faixa de frete pelo peso do pacote, em gramas:
// até 300 g (inclusive) → "carta"
// de 301 a 2000 g → "pacote leve"
// de 2001 a 30000 g → "pacote"
// acima de 30000 g → "transportadora"
// Pesos zero ou negativos → "inválido".
// Use switch sem expressão (switch true), com uma faixa por case.
// Exemplo: FreightBand(500) → "pacote leve"
func FreightBand(grams int) string {
	switch {
	case grams < 0:
		return "inválido"
	case grams <= 300:
		return "carta"
	case grams <= 2000:
		return "pacote leve"
	case grams <= 30000:
		return "pacote"
	default:
		return "transportadora"
	}
}
