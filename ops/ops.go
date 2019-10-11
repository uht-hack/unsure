package ops

import (
	"context"
	"github.com/uht-hack/unsure/db/rounds"
	"github.com/uht-hack/unsure/state"
	"strconv"
)

// JoinRound tells the player to join the specified round.
func JoinRound(ctx context.Context, s *state.State, roundID string) error {
	// Create a round
	// TODO(jonathan): Check match
	rID, err := strconv.Atoi(roundID)
	if err != nil {
		return err
	}

	err = rounds.Join(ctx, s.UhtDB().DB, "uht", 0, rID)
	if err != nil {
		return err
	}

	r, err := rounds.LookupByIndex(ctx, s.UhtDB().DB, rID)
	if err != nil {
		return err
	}

	// Ask the engine client to join the round
	cl := s.EngineClient()
	err = cl.JoinRound(ctx,"uht", player, int64(rID))
	if err != nil {
		return err
	}

	// Move the round into joined state
	err = rounds.ToJoined(ctx, s.UhtDB().DB, r.ID, rounds.RoundStatusJoin, r.CreatedAt, rounds.RoundState{})
	if err != nil {
		return err
	}

	return nil
}