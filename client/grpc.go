package client

import (
	"context"
	"github.com/corverroos/unsure"
	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"
	"github.com/uht-hack/unsure/uhtpb"
	"google.golang.org/grpc"

	uht "github.com/uht-hack/unsure"
	pb "github.com/uht-hack/unsure/uhtpb"

)

var _ uht.UhtClient = (*client)(nil)

type client struct {
	address   string
	rpcConn   *grpc.ClientConn
	rpcClient pb.UhtClient
}

func (c *client) Ping(ctx context.Context) error {
	_, err := c.rpcClient.Ping(ctx, &pb.Empty{})
	return err
}

func (c *client) Stream(ctx context.Context, after string, opts ...reflex.StreamOption) (reflex.StreamClient, error) {
	sFn := reflex.WrapStreamPB(func(ctx context.Context,
		req *reflexpb.StreamRequest) (reflex.StreamClientPB, error) {
		return c.rpcClient.Stream(ctx, req)
	})
	return sFn(ctx, after, opts...)
}

func (c *client) RoundData(ctx context.Context, round int64) (*uht.CollectRoundRes, error) {
	res, err := c.rpcClient.RoundData(ctx, &uhtpb.CollectRoundReq{
		RoundId: round,
	})
	if err != nil {
		return nil, err
	}

	var realPlayers = make([]uht.CollectPlayer, len(res.Players))
	for i, p := range res.Players {
		realPlayers[i].Name = p.Name
		realPlayers[i].Part = int(p.Part)
	}

	return &uht.CollectRoundRes {
		Rank: int(res.Rank),
		Players: realPlayers,
	}, nil
}

type option func(*client)

func WithAddress(address string) option {
	return func(c *client) {
		c.address = address
	}
}

func New(addr string, opts ...option) (*client, error) {
	c := client{
		address: addr,
	}
	for _, o := range opts {
		o(&c)
	}

	var err error
	c.rpcConn, err = unsure.NewClient(c.address)
	if err != nil {
		return nil, err
	}

	c.rpcClient = pb.NewUhtClient(c.rpcConn)

	return &c, nil
}
