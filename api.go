package unsure

import (
	"context"
	"github.com/luno/reflex"
)

type UhtClient interface {
	Ping(context.Context) error
	Stream(ctx context.Context, after string, opts ...reflex.StreamOption) (reflex.StreamClient, error)
	RoundData(ctx context.Context, round int64) (*CollectRoundRes, error)
}

type CollectRoundRes struct {
	Players []CollectPlayer
}

type CollectPlayer struct {
	Name string
	Rank    int
	Part int
}
