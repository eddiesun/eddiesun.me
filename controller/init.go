package controller

import (
	"github.com/eddiesun.me/config"
	"log"
	"net/http"
	"text/template"
)

func initialize(w http.ResponseWriter, r *http.Request, viewName string) (tmpl *template.Template, err error) {
	if err != nil {
		log.Fatalln("controller.initalize Error: ", err)
	}

	tmpl = template.New(viewName)
	tmpl = template.Must(tmpl.ParseGlob(config.ABS_PATH_TO_THIS_FILE + "view/*.html"))

	return
}
