package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"golang.org/x/net/proxy"
)

var client = createTorClient("tcp", "127.0.0.1", "9050")

func createTorClient(protocol string, address string, port string) *http.Client {
	dialer, err := proxy.SOCKS5(protocol, address+":"+port, nil, proxy.Direct)
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

func getInfoHandler(req *http.Request, writer http.ResponseWriter) (int, string) {
	currentState := State{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&currentState)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Head(currentState.Url)
	if err != nil {
		return 404, err.Error()
	}
	buffer := make([]byte, 1024)
	_, err = resp.Body.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	print(string(buffer))
	return 200, "GET INFO"
}

func main() {
	martini.Env = martini.Prod
	m := martini.Classic()
	m.Post("/links", getLinksHandler)
	m.Post("/info", getInfoHandler)
	m.Run()
}
