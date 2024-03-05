package messangerservice

import "chapar/internals/core/ports"

type MessangerService struct {
	Bridges ports.Bridges
}

func NewMessangerService(bridges ports.Bridges) *MessangerService {
	return &MessangerService{Bridges: bridges}
}
