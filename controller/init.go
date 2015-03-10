package controller

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func initialize(w http.ResponseWriter, r *http.Request, viewFile string) (tmpl *template.Template, err error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("controller.initalize Error: ", err)
	}

	templateFiles := []string{
		filepath.Join(cwd, "./view/"+viewFile),
		filepath.Join(cwd, "./view/header.tmpl"),
		filepath.Join(cwd, "./view/footer.tmpl"),
	}

	tmpl, err = template.ParseFiles(templateFiles...)
	if err != nil {
		return
	}
	return
}
