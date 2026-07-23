package main

import "net/http"

func main() {
	fileserver := http.FileServer(http.Dir("./public"))

	mux := http.NewServeMux()

	mux.Handle("/", fileserver)

	http.ListenAndServe(":8080", mux)
}
