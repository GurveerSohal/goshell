package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"

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

		cmdList := strings.Fields(string(p))

		path, err := exec.LookPath(cmdList[0])
		if err != nil {
			writeMessage(messageType, []byte("Could not run your command"), conn)
			continue
		}
		
		cmdList[0] = path
 		cmd := exec.Command(path, cmdList[1:]...)
		out, err := cmd.Output()
		if err != nil {
			writeMessage(messageType, []byte("Could not run your command"), conn)
		} else {
			writeMessage(messageType, out, conn)
		}
	}
}

func writeMessage(messageType int, p []byte, conn *websocket.Conn) {
	if err := conn.WriteMessage(messageType, p); err != nil {
		fmt.Println("error when writing message")
		return
	}
}