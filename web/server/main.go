package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}
func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/ws", handleWs)
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("settin server")
}

func handleWs(w http.ResponseWriter, r *http.Request) {
	// allow all origins (CORS)
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("error when upgrading to socket")
	}

	fmt.Println("Upgraded to websocket")
	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error when reading message")
			return
		}

		fmt.Println("received", string(p))

		if err := conn.WriteMessage(messageType, []byte("Received!")); err != nil {
			fmt.Println("error when writing message")
			return
		}
	}
}