package domain

type Hub struct {
	clients map[uint]Client
	mailman chan Message
}

func NewHub() *Hub {
	return &Hub{
		mailman: make(chan Message),
		clients: make(map[uint]Client),
	}
}

func (h *Hub) Register(user User) chan Message {
	h.clients[user.ID] = Client{reciver: user.Reciver}
	return h.mailman
}

func (h *Hub) UnRegister(user User) {
	delete(h.clients, user.ID)
}

type Client struct {
	// lock    sync.RWMutex
	reciver chan Message
}

func (h *Hub) Run() {
	for message := range h.mailman {
		h.clients[message.To].reciver <- message
	}
}
