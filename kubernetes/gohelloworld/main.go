package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	log.Println("Starting server")

	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello World")

}
