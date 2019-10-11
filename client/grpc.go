package client

import (
	"context"
	"github.com/corverroos/unsure"
	pb "github.com/corverroos/unsure/engine/enginepb"
	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"
	"google.golang.org/grpc"

	uht "github.com/uht-hack/unsure"
)

var _ uht.UhtClient = (*client)(nil)

type client struct {
	address   string
	rpcConn   *grpc.ClientConn
	rpcClient pb.EngineClient
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

	c.rpcClient = pb.NewEngineClient(c.rpcConn)

	return &c, nil
}
