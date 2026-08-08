package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Requisição iniciada")
	defer log.Println("Requisição finalizada")

	select {
	case <-time.After(4 * time.Second):
		log.Println("Requisição processada")
		w.Write([]byte("Requisição processada"))
	case <-ctx.Done():
		log.Println("Requisição cancelada")
	}
}
