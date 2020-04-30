package main

// Room maintains the set of active clients and broadcasts messages to the clients
type Room struct {
	broadcast chan []byte
	join      chan *Client
	leave     chan *Client
	clients   map[*Client]bool
}

func newRoom() *Room {
	return &Room{
		broadcast: make(chan []byte),
		join:      make(chan *Client),
		leave:     make(chan *Client),
		clients:   make(map[*Client]bool),
	}
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				close(client.send)
			}
		case msg := <-r.broadcast:
			for client := range r.clients {
				select {
				case client.send <- msg:
				default:
					close(client.send)
					delete(r.clients, client)
				}
			}
		}
	}
}
