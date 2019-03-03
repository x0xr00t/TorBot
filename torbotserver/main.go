package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
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

func getInfoHandler(req *http.Request, r render.Render) {
	currentState := new(State)
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(currentState)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Head(currentState.Url)
	if err != nil {
		r.JSON(http.StatusInternalServerError, "Error finding url")
	} else {
		r.JSON(resp.StatusCode, resp.Header)
	}
}

func main() {
	martini.Env = martini.Prod
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Post("/links", getLinksHandler)
	m.Post("/info", getInfoHandler)
	m.Run()
}
