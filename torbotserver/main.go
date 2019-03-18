package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
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

func getLinksHandler(w http.ResponseWriter, r *http.Request) {
	return
}

type State struct {
	Url    string `json::url`
	Option string `json::option,omitempty`
}

func getInfoHandler(w http.ResponseWriter, r *http.Request) {
	currentState := new(State)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(currentState)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Head(currentState.Url)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(resp.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
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
		conn.WriteMessage(websocket.TextMessage, []byte("Hello World."))
	})
	http.HandleFunc("/links", getLinksHandler)
	http.HandleFunc("/info", getInfoHandler)
	log.Print("Serving on port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
