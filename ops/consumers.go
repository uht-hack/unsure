package ops

import (
	"context"
	"database/sql"
	"github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rpatterns"
	"github.com/uht-hack/unsure"
	"github.com/uht-hack/unsure/db/cursors"
	"github.com/uht-hack/unsure/db/events"
	"github.com/uht-hack/unsure/db/rounds"
	"github.com/uht-hack/unsure/state"
	"strconv"
)

const engineEventsConsumer = "engine_updates"

// GetCTX returns a new context.
// TODO(jonathan): Add crashing functionality into this.
func GetCTX() context.Context {
	ctx := context.Background()
	return ctx
}

func consumeEngineEventsForever(s *state.State) {
	f := func(ctx context.Context, fate fate.Fate, e *reflex.Event) error {
		// Handle round join events.
		if reflex.IsAnyType(e.Type, engine.EventTypeRoundJoin) {
			return JoinRound(ctx, s, e.ForeignID)
		}

		// Handle round collect events.
		if reflex.IsAnyType(e.Type, engine.EventTypeRoundCollect) {
			return CollectRound(ctx, s, e.ForeignID)
		}

		return nil
	}

	cursorStore := cursors.ToStore(s.UhtDB().DB)
	consumable := reflex.NewConsumable(s.EngineClient().Stream, cursorStore)
	consumer := reflex.NewConsumer(engineEventsConsumer, f, reflex.WithConsumerActivityTTL(-1))

	rpatterns.ConsumeForever(GetCTX, consumable.Consume, consumer)
}

type consumerName = reflex.ConsumerName
const playerEventsConsumer consumerName = "consume_player_update"

func ConsumeAllPlayersForever(s *state.State) {
	for i := 0; i < 3; i++ {
		go ConsumePlayerEvents(s.UhtDB().DB, s.EngineClient(), s.UhtClient(i), false)
	}

	// Consume Own Events
	go ConsumePlayerEvents(s.UhtDB().DB, s.EngineClient(), nil, true)
}

func ConsumePlayerEvents(dbc *sql.DB, ec engine.Client, uhtClient unsure.UhtClient, isOwnEvents bool) {
	f := func(ctx context.Context, fate fate.Fate, e *reflex.Event) error {


		if reflex.IsAnyType(e.Type, rounds.RoundStatusCollected) || reflex.IsAnyType(e.Type, rounds.RoundStatusSubmit) {

			// Do lookup for players' data and add to current round state.

			id, _ := strconv.Atoi(e.ForeignID);
			roundData,err := uhtClient.RoundData(ctx, int64(id))
			if err !=nil {
				return err
			}
			err = AddPlayerState(ctx,dbc,*roundData,id)
			return err


		}

		if reflex.IsAnyType(e.Type, rounds.RoundStatusSubmit) {
			// Try submit
			return AttemptSubmit(ctx, dbc, ec, e.ForeignID)
		}

		return nil

	}

	cursorStore := cursors.ToStore(dbc)
	c := reflex.NewConsumer(playerEventsConsumer, f)

	stream := uhtClient.Stream
	if isOwnEvents {
		stream = events.ToStream(dbc)
	}
	consumable := reflex.NewConsumable(stream, cursorStore)
	rpatterns.ConsumeForever(context.Background, consumable.Consume, c)

}
