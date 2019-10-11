package server

import (
	"github.com/luno/reflex"
)

//var _ pb.EngineServer = (*Server)(nil)

// Server implements the engine grpc server.
type Server struct {
	b       Backends
	rserver *reflex.Server
	stream  reflex.StreamFunc
}

// New returns a new server instance.
func New(b Backends) *Server {
	return &Server{
		b:       b,
		rserver: reflex.NewServer(),
	}
}

func (srv *Server) Stop() {
	srv.rserver.Stop()
}