package main

import (
	"fmt"
	_ "github.com/eddiesun.me/controller"
	"log"
	"net/http"
)

func main() {

	// register handler
	http.HandleFunc("/", handler)

	// listen and serve
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln("ListenAndServe Error: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
