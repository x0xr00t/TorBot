package main

import (
	"github.com/go-martini/martini"
)

func getLinksHandler() string {
	return "GET LINKS"
}

func getInfoHandler() string {
	return "GET INFO"
}

func main() {
	m := martini.Classic()
	m.Get("/links", getLinksHandler)
	m.Get("/info", getInfoHandler)
	m.Run()
}
