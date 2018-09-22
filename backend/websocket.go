package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"google.golang.org/appengine/log"

	"google.golang.org/appengine"
	"encoding/json"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type Hub struct {
	clients map[*Client]bool
	register chan *Client
}

func newHub() *Hub  {
	return &Hub{
		clients:make(map[*Client]bool),
		register: make(chan *Client),
	}
}

func (hub *Hub) run()  {
	for {
		select {
			case client := <- hub.register:
				hub.clients[client] = true
		}
	}
}

type Message struct {
	number uint64
}

type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan Message
}


func initClientConn(conn *websocket.Conn) {
	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(60 * time.Second)); return nil})
}

func (client *Client) readPump() {
	defer func() {
		// client.hub.unregister <- client
		client.conn.Close()
	}()
	initClientConn(client.conn)
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {

		}
		response := Message{}
		json.Unmarshal(message, &response)
		response.number = response.number + 1
		client.send <- response
	}

}

func (client *Client) writePump() {
	for {
		select {
			case message := <- client.send:
				writer, err := client.conn.NextWriter(websocket.TextMessage)
				if err != nil {
					return
				}
				js, err := json.Marshal(message)
				if err != nil {
					return
				}
				writer.Write(js)
		}
	}
}


func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request)  {
	ctx := appengine.NewContext(r)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf(ctx, "Error upgrading connection", err)
	}

    client := &Client{hub: hub, conn: conn, send: make(chan Message)}
    client.hub.register <- client

    go client.writePump()
    go client.readPump()
}
