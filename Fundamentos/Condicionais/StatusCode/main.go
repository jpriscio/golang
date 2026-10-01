package main

// Códigos de status HTTP mais comuns do dia a dia.
// StatusText devolve a frase padrão de um código HTTP:
// 200 → "OK", 201 → "Created", 204 → "No Content",
// 301 → "Moved Permanently", 404 → "Not Found",
// 500 → "Internal Server Error".
// Para qualquer outro código, devolva "Unknown".
// Use switch, não uma cadeia de ifs.
// Exemplo: StatusText(404) → "Not Found"
func StatusText(code int) string {
	switch code {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 204:
		return "No Content"
	case 301:
		return "Moved Permanently"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown"
	}

}
