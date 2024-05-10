package main

import (
	"log"
	"net/http"
    "fmt"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
    for {
    // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
    // print out that message for clarity
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }

    }
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    // upgrade this connection to a WebSocket
    // connection
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }

    defer conn.Close() 

    log.Println("Client Connected")
    err = conn.WriteMessage(1, []byte("Hi Client!"))
    if err != nil {
        log.Println(err)
    }

    reader(conn)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./index.html")
}

func setupRoutes() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/ws", wsEndpoint)
}

func main() {
    setupRoutes()
    log.Fatal(http.ListenAndServe(":8080", nil))
}