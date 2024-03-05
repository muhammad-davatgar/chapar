package ports

import "chapar/internals/core/domain"

type Bridges interface {
	Send(message domain.Message) error
	Listener(id uint) chan domain.Message
}
