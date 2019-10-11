package ops

import (
	"context"
	"time"

	"github.com/uht-hack/unsure/db/cursors"
	"github.com/corverroos/unsure/engine"
	"github.com/luno/fate"
	"github.com/luno/reflex"
	"github.com/luno/reflex/rpatterns"
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
	consumable := reflex.NewConsumable(s.EngineClient().Stream(context.Background(), "", reflex.WithStreamFromHead()), cursorStore)
	consumer := reflex.NewConsumer(engineEventsConsumer, f, reflex.WithConsumerActivityTTL(-1))

	rpatterns.ConsumeForever(GetCTX, consumable.Consume, consumer)
}