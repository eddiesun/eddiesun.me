package controller

import (
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	log.Println("IndexController - incoming request")
	// initialize controller
	tmpl, err := initialize(w, r, "index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("controller.Index Error: ", err)
	}

	// set up data
	data := struct {
		Name string
	}{
		Name: "a testing struct",
	}

	// execute
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("controller.Index Error: ", err)
	}
}
