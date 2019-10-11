package client

import (
	"github.com/corverroos/unsure"
	pb "github.com/corverroos/unsure/engine/enginepb"
	"google.golang.org/grpc"

	uht "github.com/uht-hack/unsure"
)

var _ uht.UhtClient = (*client)(nil)

type client struct {
	address   string
	rpcConn   *grpc.ClientConn
	rpcClient pb.EngineClient
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