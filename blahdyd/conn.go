package main

import (
	"code.google.com/p/go.net/websocket"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	buff chan []byte
}

func (c *connection) reader() {
	for {
		var message []byte
		_, err := c.ws.Read(message)
		if err != nil {
			print(err)
		}
		print(message)
		h.broadcast <- message
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.buff {
		_, err := c.ws.Write([]byte(message))
		if err != nil {
			print(err)
			break
		}
	}
	c.ws.Close()
}

func wsHandler(ws *websocket.Conn) {
	c := &connection{buff: make(chan []byte, 256), ws: ws}
	n, err := ws.Write([]byte("Sad"))
	if err != nil {
		print(err)
	}
	print(n)
	err = websocket.Message.Send(ws, "But")
	if err != nil {
		print(err)
	}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}
