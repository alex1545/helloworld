package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

func main() {
	port := ":8080"
	http.HandleFunc("/", helloHandler)
	fmt.Println("Listening at", port[1:])
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
