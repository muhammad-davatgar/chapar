package http

import (
	"chapar/internals/core/ports"
	"chapar/internals/core/services/auth"
)

type HTTPServer struct {
	innerBridges   ports.InnerBridges
	authentication auth.AuthenticationService
}

func NewHttpService(bridges ports.InnerBridges, authenctication auth.AuthenticationService) HTTPServer {
	return HTTPServer{
		innerBridges:   bridges,
		authentication: authenctication,
	}
}
