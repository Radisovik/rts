package main

import (
	"embed"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"io/fs"
	"log"
	"net/http"
	"server/messages"
)

//go:embed public/*
var webRoot embed.FS
var upgrader = websocket.Upgrader{} // use default options

func main() {
	htmlContent, err := fs.Sub(webRoot, "public")
	chk(err)
	rootFileServer := http.FileServer(http.FS(htmlContent))
	http.Handle("/", rootFileServer)
	http.HandleFunc("/ws", gameConnection)
	log.Printf("Staring web server")
	err = http.ListenAndServe(":8000", nil)
	chk(err)
}

func gameConnection(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade to a websocket: %v", err)
		return
	}
	defer c.Close()
	log.Printf("New inbound websocket connection from %v", r.RemoteAddr)
	fs := &messages.FromServer{
		Type:  messages.FromServer_Hello,
		Hello: &messages.Hello{Greeting: "hello world"},
	}
	marshal, err := proto.Marshal(fs)
	chk(err)
	err = c.WriteMessage(websocket.BinaryMessage, marshal)
	chk(err)
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("Error while reading websocket msg")
			break
		}
		if mt != websocket.BinaryMessage {
			log.Printf("Someone sent an ASCII message.. we don't do that")
			break
		}
		log.Printf("Here is our cool message: %v", message)
	}
}

func chk(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
