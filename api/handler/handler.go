package handler

import (
	"google.golang.org/grpc"
)

type Server struct {
	Usermanagement      *grpc.ClientConn
	Gargardenmanagement *grpc.ClientConn
	Sustainability      *grpc.ClientConn
	Community           *grpc.ClientConn
}

type HandlerConfig struct {
}

func NewHandlerConfig(conn *Server) *HandlerConfig {
	return &HandlerConfig{}
}
