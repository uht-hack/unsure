package ops

import (
	"context"
	"github.com/uht-hack/unsure/db/rounds"
	"github.com/uht-hack/unsure/state"
)

// JoinRound tells the player to join the specified round.
func JoinRound(ctx context.Context, s *state.State, roundID string) error {
	// Create a round
	rounds.Join(ctx, s.UhtDB(), "uht")

	// Ask the engine client to join the round

	// Move the round into joined state

}