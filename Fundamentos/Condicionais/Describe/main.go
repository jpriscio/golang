package main

import "fmt"

// Describe recebe um valor de qualquer tipo e devolve uma descrição:
// int → "inteiro: <valor>" (use fmt.Sprintf com %d)
// string → "texto de <n> bytes" (n = len do valor)
// bool → "booleano: <valor>" (use %t)
// qualquer outro tipo → "tipo não suportado"
// Use um type switch: switch v := value.(type) { ... }.
// Exemplo: Describe(42) → "inteiro: 42"
// Exemplo: Describe("oi") → "texto de 2 bytes"
func Describe(value any) string {

	switch v := value.(type) {
	case int:
		return fmt.Sprintf("inteiro: %d", v)
	case string:
		return fmt.Sprintf("texto de %d bytes", len(v))
	case bool:
		return fmt.Sprintf("booleanos: %t", v)
	default:
		return "tipo não suportado"
	}
}
