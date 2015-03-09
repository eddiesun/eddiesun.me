package main

type (
	Route struct {
		Url        string
		Controller string
		View       string
	}
)

var (
	Routes = []Route{
		{"/", "index", "index.html"},
	}
)
