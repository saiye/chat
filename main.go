package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

// Message represents a message sent from a client to the server.
type Message struct {
	From    string `json:"from"`
	Message string `json:"message"`
}

var (
	clients     = make(map[*websocket.Conn]bool)
	broadcaster = make(chan Message)
	clientsMu   sync.Mutex
)

// Client represents a single websocket client.
type Client struct {
	conn *websocket.Conn
	send chan Message
}

// readPump pumps messages from the websocket connection to the broadcaster.
func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
		clientsMu.Lock()
		delete(clients, c.conn)
		clientsMu.Unlock()
	}()
	for {
		var msg Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		broadcaster <- msg
	}
}

// writePump pumps messages from the broadcaster to the websocket connection.
func (c *Client) writePump() {
	ticker := time.NewTicker(time.Second * 5)
	pongWait := time.Millisecond * 500
	defer func() {
		ticker.Stop()
		close(c.send)
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// The broadcaster channel has been closed.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.conn.WriteJSON(message); err != nil {
				fmt.Sprintf("WriteMessageError:%v", err)
				return
			}
		case <-ticker.C:
			if err := c.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(pongWait)); err != nil {
				fmt.Sprintf("PingMessageError:%v", err)
				return
			}
		}
	}
}

var up = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//allowedOrigin := "https://yourfrontend.com"
		//return r.Header.Get("Origin") == allowedOrigin
		return true
	},
}

// serveWs handles websocket requests from clients.
func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	client := &Client{conn: conn, send: make(chan Message, 256)}
	clientsMu.Lock()
	clients[conn] = true
	clientsMu.Unlock()
	go client.readPump()
	go client.writePump()
}

func broadcasterHandler() {
	for {
		select {
		case message, ok := <-broadcaster:
			if !ok {
				// The broadcaster channel has been closed.
				return
			}
			clientsMu.Lock()
			for con, flag := range clients {
				if flag {
					con.WriteJSON(message)
				}
			}
			clientsMu.Unlock()
		}
	}
}

func main() {
	//{"form":"buffer","message":"run go"}
	//{"form":"xiuluo","message":"run xiuluo"}
	go broadcasterHandler()

	http.HandleFunc("/ws", serveWs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
