package matches

import (
	"database/sql"

	"github.com/luno/reflex"
	"github.com/luno/reflex/rsql"
)

var events = rsql.NewEventsTableInt("match_events")

// ToStream returns a reflex stream for Match events.
func ToStream(dbc *sql.DB) reflex.StreamFunc {
	return events.ToStream(dbc)
}
