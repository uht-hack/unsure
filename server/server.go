package server

import (
	"context"
	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"
	"github.com/uht-hack/unsure/db/events"
	"github.com/uht-hack/unsure/db/rounds"
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

func (srv *Server) RoundData(ctx context.Context, req *uhtpb.CollectRoundReq) (*uhtpb.CollectRoundRes, error) {
	round, err := rounds.Lookup(ctx, srv.b.UhtDB().DB, req.RoundId)
	if err != nil {
		return nil, err
	}

	players := make([]*uhtpb.CollectPlayer, len(round.State.Players))
	for i, p := range round.State.Players {
		players[i] = &uhtpb.CollectPlayer{
			Name: p.Name,
			Rank: int64(p.Rank),
			Parts: p.Parts,
		}
	}

	return &uhtpb.CollectRoundRes{
		Players: players,
	}, nil
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
