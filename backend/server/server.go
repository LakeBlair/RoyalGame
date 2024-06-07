package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/LakeBlair/royalgame/backend/internal"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
    upgrader  websocket.Upgrader
}

var (
    sessions = make(map[string]*internal.GameSession)
)

func NewWSHandler(upgrader websocket.Upgrader) *WebSocketHandler {
    return &WebSocketHandler{
        upgrader:  upgrader,
    }
}

func (s *WebSocketHandler) createSession(w http.ResponseWriter, r *http.Request) {
    log.Println("Creating New Session")
    sessionID := uuid.New().String() // Generate a unique session ID
    sessions[sessionID] = &internal.GameSession{
        ID: sessionID,
        MoveDataChannel: make(chan int),
    }
    w.Write([]byte(sessionID))
}

func (s *WebSocketHandler) play(w http.ResponseWriter, r *http.Request) {
    conn, conn_err := s.upgrader.Upgrade(w, r, nil)
    if conn_err != nil {
        log.Println("WebSocket Upgrade Error:", conn_err)
        return
    }
    // defer conn.Close()

    sessionID := r.URL.Query().Get("session_id")
    if sessionID == "" {
        http.Error(w, "Session ID required", http.StatusBadRequest)
        return
    }

    // Add new client connection to session
    session, exists := sessions[sessionID]

    if !exists {
        log.Println("Session doesn't exist, exiting...")
        return
    }
    if len(session.Connections) < 2 {
        session.Connections = append(session.Connections, conn)
        sessions[sessionID] = session
    }

    _, message, read_err := conn.ReadMessage()
    if read_err != nil {
        log.Println("Read Error:", read_err)
        return
    }
    log.Printf("Received: %s", message)

    var Msg internal.GameMessage

    if err := json.Unmarshal(message, &Msg); err != nil {
        log.Println("Data Error:", err)
        return
    }

    switch Msg.Msg_Type {
    case "Start":
        if len(session.Connections) == 2 {
            go internal.Play(session)
        }
    case "Move":
        log.Printf("Move: %d", Msg.Move)
        session.MoveDataChannel <- Msg.Move
    }
}

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hellow World")
}

func (s *WebSocketHandler) setupRoutes() {
    log.Println("Setting up routes")
    buildDir := http.Dir("frontend/build")
    fs := http.FileServer(buildDir)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fs.ServeHTTP(w, r)
    })
    http.HandleFunc("/play", s.play)
    http.HandleFunc("/create-session", s.createSession)
}


func (s *WebSocketHandler) LaunchGame() {
    s.setupRoutes()

    port := os.Getenv("PORT")
    if port == "" {
        log.Println("$PORT set to 8080")
        port = "8080"
    }

    log.Fatal(http.ListenAndServe(":" + port, nil))
}