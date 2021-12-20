package main

import (
	"embed"
	"fmt"
	proto "github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed public/*
var webRoot embed.FS
var upgrader = websocket.Upgrader{} // use default options

func main() {
	setupWebHandler()

	http.HandleFunc("/ws", gameConnection)
	log.Printf("Staring web server")
	err := http.ListenAndServe(":8000", nil)
	chk(err)
}

func setupWebHandler() {
	if exists("public") {
		log.Printf("Noticed public directory -- will serve that!")
		fs := http.FileServer(http.Dir("./public"))
		http.Handle("/", fs)
	} else {
		log.Printf("No public directory found -- will serve embedded files")
		htmlContent, err := fs.Sub(webRoot, "public")
		chk(err)
		rootFileServer := http.FileServer(http.FS(htmlContent))
		http.Handle("/", rootFileServer)
	}

}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func gameConnection(w http.ResponseWriter, r *http.Request) {
	var currentUser = ""
	var info = func(f string, args ...any) {
		msg := fmt.Sprintf(f, args...)
		log.Printf("[%v] %s", currentUser, msg)
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		info("Failed to upgrade to a websocket: %v", err)
		return
	}
	defer c.Close()
	defer func() {
		info("Closing")
	}()
	currentUser = r.RemoteAddr
	info("New inbound websocket connection from %v", currentUser)
	fs := &FromServer{
		Epoch: uint64(time.Now().Unix()),
		Msg:   &FromServer_Hello{Hello: &Hello{Greeting: "Hello from Server"}},
	}

	marshal, err := proto.Marshal(fs)
	chk(err)
	err = c.WriteMessage(websocket.BinaryMessage, marshal)
	chk(err)
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			info("Error while reading websocket msg")
			break
		}
		if mt != websocket.BinaryMessage {
			info("Someone sent an ASCII message.. we don't do that")
			break
		}
		fc := FromClient{}
		err = proto.Unmarshal(message, &fc)
		if err != nil {
			info(err.Error())
			return
		}
		info("Here is our cool message: %v", fc.GetHello())

	}
}

func chk(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
