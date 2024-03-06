package ports

import "chapar/internals/core/domain"

type InnerBridges interface {
	Register(domain.User) chan domain.Message
	UnRegister(domain.User)
}
