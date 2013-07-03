package main

import (
	"log"
	"time"
)

type hub struct {
	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan string

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

var h = hub{
	broadcast:   make(chan string, 100),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			log.Println("Adding client")
			h.connections[c] = true
		case c := <-h.unregister:
			log.Println("Removing client")
			delete(h.connections, c)
			close(c.send)
		case m := <-h.broadcast:
			broadcast(h, m)
		}
	}
}

func broadcast(h *hub, m string) {
	conns := make(map[*connection]bool)
	for c := range h.connections {
		conns[c] = true
	}

	timeout := time.After(5 * time.Second)
	timedOut := false
	for !timedOut && len(conns) > 0 {
		for c := range conns {
			select {
			case <-timeout:
				timedOut = true
				break
			case c.send <- m:
				delete(conns, c)
			default:
				continue
			}
		}
	}

	for c := range conns {
		delete(h.connections, c)
		close(c.send)
		go c.ws.Close()
	}
}
