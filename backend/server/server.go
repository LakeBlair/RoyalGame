package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type ApiResponse struct {
    Message string `json:"message"`
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
    err = conn.WriteMessage(websocket.TextMessage, []byte("Hi Client!"))
    if err != nil {
        log.Println(err)
    }

    reader(conn)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
    response := ApiResponse{Message: "Hello from the Golang API!"}
    log.Println("Hello from the Golang API!")
    json.NewEncoder(w).Encode(response)
}


// func homePage(w http.ResponseWriter, r *http.Request) {
//     http.ServeFile(w, r, "../../frontend/build")
// }

func setupRoutes() {
    log.Println("Setting up routes")
    buildDir := http.Dir("../../frontend/build")
    fs := http.FileServer(buildDir)
    http.Handle("/", fs)
    http.HandleFunc("/ws", wsEndpoint)
    http.HandleFunc("/api/data", dataHandler)
}

func LaunchGameServer() {
    setupRoutes()
    log.Fatal(http.ListenAndServe(":8080", nil))
}