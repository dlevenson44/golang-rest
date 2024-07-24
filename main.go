package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Routes
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)
	fmt.Println(err)
}

// Handler functions-- similar to Controller functions
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got / request\n")
	io.WriteString(w, "This is my site!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
