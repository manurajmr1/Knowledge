package main

import (
	
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/favicon.ico", handlerICon)
	http.HandleFunc("/", home)               // set router
	err := http.ListenAndServe(":9002", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func home(w http.ResponseWriter, re *http.Request) {
	//val := 20 * 1200 / 12 * 22
	//fmt.Fprintf(w, "Hello Gopher%d", val) // send data to client side
}

func handlerICon(w http.ResponseWriter, r *http.Request) {}
