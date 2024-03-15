package ports

import (
	"chapar/internals/core/domain"
)

type MessangerServices interface {
	GetChat(first, second uint) ([]domain.Message, error)
	Send(domain.Message)
	Listen() chan domain.Message
}

type UserDB interface {
	Create(domain.User) (domain.User, error)
	GetUser(domain.Credentials) (domain.User, error)
}
