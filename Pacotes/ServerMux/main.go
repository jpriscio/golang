package main

import "net/http"

func Homeland(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Homeland)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello tmb"))
	})

	http.ListenAndServe(":8080", mux)

}
