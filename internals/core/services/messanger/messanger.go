package messangerservice

import (
	"chapar/internals/core/domain"
	"chapar/internals/core/ports"
)

type MessangerService struct {
	bridges ports.InnerBridges
}

func NewMessangerService() *MessangerService {
	hub := domain.NewHub()
	go hub.Run()
	return &MessangerService{bridges: hub}
}

func (s *MessangerService) Register(user domain.User) chan domain.Message {
	return s.bridges.Register(user)
}
func (s *MessangerService) UnRegister(user domain.User) {
	s.bridges.UnRegister(user)
}
