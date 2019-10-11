package server

import (
	"context"
	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"
	"github.com/uht-hack/unsure/db/events"
	"github.com/uht-hack/unsure/uhtpb"
)

var _ uhtpb.UhtServer = (*Server)(nil)

// Server implements the engine grpc server.
type Server struct {
	b       Backends
	rserver *reflex.Server
	stream  reflex.StreamFunc
}

func (srv *Server) Ping(ctx context.Context, req *uhtpb.Empty) (*uhtpb.Empty, error) {
	return req, nil
}

func (srv *Server) Stream(req *reflexpb.StreamRequest, ss uhtpb.Uht_StreamServer) error {
	return srv.rserver.Stream(srv.stream, req, ss)
}

// New returns a new server instance.
func New(b Backends) *Server {
	return &Server{
		b:       b,
		rserver: reflex.NewServer(),
		stream:  events.ToStream(b.UhtDB().DB),
	}
}

func (srv *Server) Stop() {
	srv.rserver.Stop()
}
