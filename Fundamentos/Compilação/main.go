package main

import "fmt"

// BuildCommand monta o comando de cross-compilation para um alvo.
// Exemplo: BuildCommand("linux", "amd64", "servidor") →
// "GOOS=linux GOARCH=amd64 go build -o servidor"
// Se output vier vazio, devolva o comando sem a parte "-o":
// BuildCommand("windows", "arm64", "") → "GOOS=windows GOARCH=arm64 go build"
func BuildCommand(goos string, goarch string, output string) string {
	// implemente:
	comando := fmt.Sprintf("GOOS=%s GOARCH=%s go build", goos, goarch)
	if output == "" {
		return comando
	}
	return comando + " -o " + output
}
