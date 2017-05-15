package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/favicon.ico", handlerICon)
	http.HandleFunc("/", home)               // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func home(w http.ResponseWriter, re *http.Request) {

	fmt.Fprintf(w, "Hello Gopher") // send data to client side
}

func handlerICon(w http.ResponseWriter, r *http.Request) {}
