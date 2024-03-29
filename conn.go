package main

import (
  "code.google.com/p/go.net/websocket"
)

type connection struct {
  // The websocket connection.
  ws *websocket.Conn

  // Buffered channel of outbound messages.
  send chan string
}

func (c *connection) reader() {
  for {
    var message string
    err := websocket.Message.Receive(c.ws, &message)
    if err != nil {
      break
    }
    h.broadcast <- message
  }
  c.ws.Close()
}

func (c *connection) writer() {
  for message := range c.send {
    err := websocket.Message.Send(c.ws, message)
    if err != nil {
      break
    }
  }
  c.ws.Close()
}

func wsHandler(ws *websocket.Conn) {
  c := &connection{send: make(chan string, 256), ws: ws}
  h.register <- c
  defer func() { h.unregister <- c }()
  go c.writer()
  c.reader()
}