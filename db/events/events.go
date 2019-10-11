package events

import (
	"context"
	"database/sql"

	"github.com/luno/reflex"
	"github.com/luno/reflex/rsql"
)

var events = rsql.NewEventsTableInt("events")

// ToStream returns a reflex stream for Match events.
func ToStream(dbc *sql.DB) reflex.StreamFunc {
	return events.ToStream(dbc)
}

// Insert inserts a reflex event into the events table
// and returns a notify function or error.
func Insert(ctx context.Context, tx *sql.Tx, foreignID int64,
	typ reflex.EventType) (func(), error) {
	return events.Insert(ctx, tx, foreignID, typ)
}

func GetTable() rsql.EventsTableInt {
	return events
}
