package ops

import (
	"context"
	"errors"
	"github.com/uht-hack/unsure/db/rounds"
	"github.com/uht-hack/unsure/state"
	"strconv"
)

func CollectRound(ctx context.Context, s *state.State, roundID string) error {
	rID, err := strconv.Atoi(roundID)
	if err != nil {
		return err
	}

	r, err := rounds.LookupByIndex(ctx, s.UhtDB().DB, rID)
	if err != nil {
		return err
	}

	// Move to collecting
	err = rounds.ToCollect(ctx, s.UhtDB().DB, r.ID, rounds.RoundStatusJoined, r.UpdatedAt, rounds.RoundState{})
	if err != nil {
		return err
	}

	// Get parts from client
	cl := s.EngineClient()
	playerState, err = cl.CollectRound(ctx, "uht", player, int64(rID))
	if err != nil {
		return err
	}

	roundState := make([]rounds.RoundPlayerState, 4)
	roundState = append(roundState, rounds.RoundPlayerState{
		Name:      *player,
		Rank:      playerState.Rank,
		Parts:     nil,
		Included:  true,
		Collected: false,
		Submitted: false,
	})

	// Move to collected
	err = rounds.ToCollected(ctx, s.UhtDB().DB, r.ID, rounds.RoundStatusJoin, r.UpdatedAt, rounds.RoundState{
		Players:roundState,
	})
	if err != nil {
		return err
	}

}

// JoinRound tells the player to join the specified round.
func JoinRound(ctx context.Context, s *state.State, roundID string) error {
	// Create a round
	rID, err := strconv.Atoi(roundID)
	if err != nil {
		return err
	}

	// TODO(jonathan): Check match
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
	ok, err := cl.JoinRound(ctx,"uht", *player, int64(rID))
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("can't join")
	}

	// Move the round into joined state
	err = rounds.ToJoined(ctx, s.UhtDB().DB, r.ID, rounds.RoundStatusJoin, r.UpdatedAt, rounds.RoundState{})
	if err != nil {
		return err
	}

	return nil
}