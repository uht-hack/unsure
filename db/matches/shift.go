package matches

import (
	"github.com/luno/shift"
	"github.com/uht-hack/unsure/db/events"
)

//go:generate shiftgen -inserter=startReq -updaters=endReq -table=matches

//                 | --> Success
// New --> Started |
//                 L --> Failed

var fsm = shift.NewFSM(events.GetTable()).
	Insert(MatchStatusStarted, startReq{}, MatchStatusEnded).
	Update(MatchStatusEnded, endReq{}).
	Build()

type startReq struct {
	Players int
}

type endReq struct {
	ID      int64
}
