package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/gorilla/websocket"
	"github.com/martini-contrib/render"
	"golang.org/x/net/proxy"
)

var client = createTorClient("tcp", "127.0.0.1", "9050")
var upgrader = websocket.Upgrader{}

type Message struct {
	Link   string `json::"link"`
	Status bool   `json::"status"`
}

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

func getLinksHandler(w http.ResponseWriter, r *http.Request) (int, string) {
	return http.StatusOK, "GET LINKS"
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

func handleMessages() {

}

func main() {
	http.HandleFunc("/test/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatalf("Error: %+v", err)
		}
		defer conn.Close()
		log.Printf("Succesfully upgraded websocket connection. Connection: %+v.\n", conn)
	})
	go http.ListenAndServe(":8080", nil)
	martini.Env = martini.Prod
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Post("/links", getLinksHandler)
	m.Post("/info", getInfoHandler)
	m.Run()
}
