package ops

import (
	"context"

	"github.com/uht-hack/unsure/db/cursors"
	"github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rpatterns"
	"github.com/uht-hack/unsure/state"
)

const engineEventsConsumer = "engine_updates"

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

	rpatterns.ConsumeForever(s.GetLeaderState().WaitUntilLeader, consumable.Consume, consumer)
}