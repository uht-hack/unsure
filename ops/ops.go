package ops

import (
	"context"
	"strconv"

	"github.com/uht-hack/unsure/db/rounds"
	"github.com/uht-hack/unsure/state"
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

	// Check if the player is included in the round
	if !r.State.Included(*player) {
		return nil
	}

	// Move to collecting
	err = rounds.ToCollect(ctx, s.UhtDB().DB, r.ID, rounds.RoundStatusJoined, r.UpdatedAt, r.State)
	if err != nil {
		return err
	}

	// Get parts from client
	cl := s.EngineClient()
	playerState, err = cl.CollectRound(ctx, "uht", player, int64(rID))
	if err != nil {
		return err
	}

	parts := map[string]int32{}
	for _, p := range playerState.Players {
		parts[p.Name] = int32(p.Part)
	}

	roundState := make([]rounds.RoundPlayerState, 4)
	roundState = append(roundState, rounds.RoundPlayerState{
		Name:      *player,
		Rank:      playerState.Rank,
		Parts:     parts,
		Included:  true,
		Collected: false,
		Submitted: false,
	})

	// Move to collected
	err = rounds.ToCollected(ctx, s.UhtDB().DB, r.ID, rounds.RoundStatusCollect, r.UpdatedAt, rounds.RoundState{
		Players:roundState,
	})
	if err != nil {
		return err
	}

	return nil
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
	included, err := cl.JoinRound(ctx,"uht", *player, int64(rID))
	if err != nil {
		return err
	}

	roundState := make([]rounds.RoundPlayerState, 4)
	roundState = append(roundState, rounds.RoundPlayerState{
		Name:      *player,
		Rank:      playerState.Rank,
		Parts:     nil,
		Included:  included,
		Collected: false,
		Submitted: false,
	})

	// Move the round into joined state
	err = rounds.ToJoined(ctx, s.UhtDB().DB, r.ID, rounds.RoundStatusJoin, r.UpdatedAt, rounds.RoundState{})
	if err != nil {
		return err
	}

	return nil
}