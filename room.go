package main

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
				client.send <- msg
			}
		}
	}
}
