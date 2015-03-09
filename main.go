package main

import (
	_ "github.com/eddiesun.me/controller"
	"html/template"
	"log"
	"net/http"
)

func main() {

	// register handler
	regRoutes()
	regStatic()

	// listen and serve
	log.Println("Application is listening...")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatalln("ListenAndServe Error: ", err)
	}
}

func regRoutes() {
	// "Routes" is defined in routes.go
	for _, route := range Routes {
		log.Println("Add Route: ", route)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

		data := struct {
			Name string
		}{
			Name: "a testing struct",
		}

		t, err := template.ParseFiles("view/index.html")
		if err != nil {
			log.Fatalln("regRoutes error: ", err)
		}
		t.Execute(w, data)
	})

}

func regStatic() {
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:]) // [1:] trim the leading '/' char
	})
}
