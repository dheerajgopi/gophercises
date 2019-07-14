package main

import (
	"log"
	"net/http"
)

func helloFunc(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloFunc)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
