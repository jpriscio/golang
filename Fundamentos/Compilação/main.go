package main

import "fmt"

func BuildCommand(goos string, goarch string, output string) string {
	// implemente:
	comando := fmt.Sprintf("GOOS=%s GOARCH=%s go build", goos, goarch)
	if output == "" {
		return comando
	}
	return comando + " -o " + output
}
