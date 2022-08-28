package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "hello!")
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/post", formHandler)

	fmt.Printf("Server started on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
