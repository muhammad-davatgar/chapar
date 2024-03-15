package ports

import "chapar/internals/core/domain"

type InnerBridges interface {
	Register(domain.HubUser) chan domain.Message
	UnRegister(domain.HubUser)
}
