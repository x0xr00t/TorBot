package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"golang.org/x/net/proxy"
)

var client = createTorClient("tcp", "127.0.0.1:9050")

func createTorClient(protocol string, address string) *http.Client {
	dialer, err := proxy.SOCKS5(protocol, address, nil, proxy.Direct)
	if err != nil {
		log.Fatal(err)
	}

	tr := &http.Transport{
		Dial: dialer.Dial,
	}
	return &http.Client{Transport: tr}
}

func getLinksHandler(req *http.Request) (int, string) {
	return 200, "GET LINKS"
}

type State struct {
	Url    string `json::url`
	Option string `json::option,omitempty`
}

func getInfoHandler(req *http.Request) (int, string) {
	currentState := State{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&currentState)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", currentState)
	return 200, "GET INFO"
}

func main() {
	martini.Env = martini.Prod
	m := martini.Classic()
	m.Post("/links", getLinksHandler)
	m.Post("/info", getInfoHandler)
	m.Run()
}
