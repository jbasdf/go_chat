package main

// import (
//   ""
// )

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
			h.connections[c] = true
		case c := <-h.unregister:
			delete(h.connections, c)
			close(c.send)
		case m := <-h.broadcast:

      for c := range h.connections {
				// for sent := c.send <- m; !sent{
			}
				select {
				case c.send <- m:
					// Remove from queue
				default:
					delete(h.connections, c)
					close(c.send)
					go c.ws.Close()
				}
			}


		}
	}
}
