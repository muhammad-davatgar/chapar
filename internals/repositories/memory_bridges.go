package repositories

import (
	"chapar/internals/core/domain"
	"chapar/shared/apperrors"
)

type MemBridges struct {
	store map[uint][]chan domain.Message
}

func (m *MemBridges) Send(message domain.Message) error {
	channels, ok := m.store[message.To]
	if !ok {
		return apperrors.ListenerNotFound
	}

	for i := range channels {
		channels[i] <- message
	}

	return nil
}

func (m *MemBridges) Listener(id uint) chan domain.Message {
	channels, ok := m.store[id]
	if !ok {
		channels = []chan domain.Message{make(chan domain.Message, 10)}
		m.store[id] = channels
		return channels[0]
	}

	channels = append(channels, make(chan domain.Message, 10))
	return channels[len(channels)-1]
}
