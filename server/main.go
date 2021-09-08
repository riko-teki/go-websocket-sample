package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("go-websocket-sample listening")
	http.HandleFunc("/ws", EchoServer)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func EchoServer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("client: %v!\n", string(p))

		time.Sleep(time.Second * 1)

		fmt.Printf("server: pong!\n")
		err = conn.WriteMessage(messageType, []byte("pong!"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
}
