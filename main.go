package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/shorten", shortenUrl)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
