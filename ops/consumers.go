package ops

import (
	"context"
	"database/sql"
	"github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rpatterns"
	"github.com/uht-hack/unsure/db/cursors"
	"github.com/uht-hack/unsure/db/events"
	"github.com/uht-hack/unsure/db/rounds"
	"github.com/uht-hack/unsure/state"
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
		if !reflex.IsAnyType(e.Type, engine.EventTypeRoundJoin) {
			return nil
		}

		return JoinRound(ctx, s, e.ForeignID)
	}

	cursorStore := cursors.ToStore(s.UhtDB().DB)
	consumable := reflex.NewConsumable(s.EngineClient().Stream(context.Background(), "", reflex.WithStreamFromHead()), cursorStore)
	consumer := reflex.NewConsumer(engineEventsConsumer, f, reflex.WithConsumerActivityTTL(-1))

	rpatterns.ConsumeForever(GetCTX, consumable.Consume, consumer)
}

type consumerName = reflex.ConsumerName
const playerEventsConsumer consumerName = "consume_player_update"

func ConsumeAllPlayersForever(s *state.State) {
	for i := 0 ;i < 3 ;i ++  {
		go ConsumePlayerEvents(s.UhtDB().DB,s.UhtClient(i).Stream, false)
	}

	// Consume Own Events
	go ConsumePlayerEvents(s.UhtDB().DB, events.ToStream(s.UhtDB().DB), true)
}

func ConsumePlayerEvents(dbc *sql.DB, stream reflex.StreamFunc, isOwnEvents bool) {



	f := func(ctx context.Context, fate fate.Fate, e *reflex.Event) error {


		if reflex.IsAnyType(e.Type, rounds.RoundStatusCollected) {

			// Do lookup for players' data

		}

		return nil

	}

	cursorStore := cursors.ToStore(dbc)
	c := reflex.NewConsumer(playerEventsConsumer, f)
	consumable := reflex.NewConsumable(stream, cursorStore)
	rpatterns.ConsumeForever(context.Background, consumable.Consume, c)

}
