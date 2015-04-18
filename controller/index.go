package controller

import (
	"github.com/eddiesun.me/config"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	log.Println("IndexController - incoming request")

	// initialize controller to get a view template
	tmpl, err := initialize(w, r, "index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("controller.Index Error: ", err)
	}

	// set up data
	data := struct {
		GoogleAnalyticsTrackingId string
	}{
		config.GOOGLE_ANALYTICS_ID,
	}

	// execute
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("controller.Index Error: ", err)
	}
}
