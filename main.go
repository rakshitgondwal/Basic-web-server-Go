package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	response, error := http.Get("https://localhost:3000")
	if error != nil {
		log.Print(error)
	}
	fmt.Print(response)
}
