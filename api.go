package unsure

import (
	"context"
	"github.com/luno/reflex"
)

type UhtClient interface {
	Ping(context.Context) error
	Stream(ctx context.Context, after string, opts ...reflex.StreamOption) (reflex.StreamClient, error)
}