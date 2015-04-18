package main

import (
	"log"
	"net/http"
)

func main() {
	// register handler
	regRoutes()
	regStatic()

	// listen and serve
	log.Println("Application is listening...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln("ListenAndServe Error: ", err)
	}
}

func regRoutes() {
	// "Routes" is defined in routes.go
	for _, route := range Routes {
		log.Printf("Add Route: %+v", route)
		http.HandleFunc(route.Pattern, route.Controller)
	}
}

func regStatic() {
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:]) // [1:] trim the leading '/' char
	})
}
