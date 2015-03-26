package main

import (
	"github.com/eddiesun.me/controller"
	"net/http"
)

type (
	Route struct {
		Name       string
		Pattern    string
		Controller func(http.ResponseWriter, *http.Request)
	}
)

var (
	Routes = []Route{
		{"Index", "/", controller.Index},
		{"ContactForm", "/contactForm", controller.ContactForm},
	}
)
